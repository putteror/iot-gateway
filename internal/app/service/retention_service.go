package service

import (
	"fmt"
	"log"
	"time"

	"github.com/putteror/iot-gateway/internal/adapter"
	"github.com/putteror/iot-gateway/internal/adapter/push"
	"github.com/putteror/iot-gateway/internal/config"
)

type RetentionService interface {
	RetentionGetPM25AndPushToDestination() error
	RetentionGetWaterSensorAndPushToDestination() error
}

// set if WEBHOOK_SERVICE_NAME = "zyta"

type RetentionServiceImpl struct {
	DestinationAdapters map[string]push.PushDataAdapter
	PM25Fetcher         adapter.GetPM25
}

func NewRetentionService(
	pm25Fetcher adapter.GetPM25,
	zytaAdapter push.ZytaPushDataAdapter,
	centAccessAdapter push.CentAccessPushDataAdapter,
) RetentionService {
	return &RetentionServiceImpl{
		DestinationAdapters: map[string]push.PushDataAdapter{
			"zyta":        zytaAdapter,       // Key คือ "zyta"
			"cent-access": centAccessAdapter, // Key คือ "cent-access"
		},
		PM25Fetcher: pm25Fetcher,
	}
}

func (s *RetentionServiceImpl) RetentionGetPM25AndPushToDestination() error {

	// 1. ตรวจสอบ Adapter และ Fetcher ที่จำเป็น
	adapter, ok := s.DestinationAdapters[config.DESTINATION_TYPE]
	if !ok {
		message := fmt.Sprintf("Error: Third party adapter '%s' not supported.", config.DESTINATION_TYPE)
		fmt.Println(message)
		return nil
	}

	// 2. เริ่ม Goroutine
	go func() {
		log.Println("Background PM2.5 fetching started...")

		for {
			// 1. ดึงข้อมูล
			pm25Value, pm10Value, humidityValue, temperatureValue, err := s.PM25Fetcher.GetPM25Value()
			if err != nil {
				log.Printf("❌ ERROR: Failed to get PM2.5 value: %v. Retrying in 10s...", err)
				time.Sleep(10 * time.Second)
				continue
			}

			adapter.PushPM25Value(pm25Value, pm10Value, humidityValue, temperatureValue)
			time.Sleep(5 * time.Minute)
		}

	}()

	return nil
}

func (s *RetentionServiceImpl) RetentionGetWaterSensorAndPushToDestination() error {
	adapter, ok := s.DestinationAdapters[config.DESTINATION_TYPE]
	if !ok {
		message := fmt.Sprintf("Error: Third party adapter '%s' not supported.", config.DESTINATION_TYPE)
		fmt.Println(message)
		return nil
	}
	go func() {
		log.Println("Background PM2.5 fetching started...")

		for {
			adapter.PushWaterSensorData()
			time.Sleep(5 * time.Minute)
		}

	}()

	return nil
}
