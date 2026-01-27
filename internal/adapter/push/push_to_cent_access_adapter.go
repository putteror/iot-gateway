package push

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/putteror/iot-gateway/internal/app/schema"
	"github.com/putteror/iot-gateway/internal/config"
)

type CentAccessPushDataAdapter interface {
	PushFaceRecognitionEventData(payload *schema.FaceRecognitionEventSchema) error
	PushEmergencyAlarmEventData(payload *schema.EmergencyAlarmEventSchema) error
	PushPM25Value(pm25Value float64, pm10Value float64, humidityValue float64, temperatureValue float64) error
	PushWaterSensorData() error
}

type CentAccessPushDataAdapterImpl struct {
}

func NewCentAccessPushDataServiceImpl() CentAccessPushDataAdapter {
	return &CentAccessPushDataAdapterImpl{}
}

type CentAccessFaceRecognitionEventSchema struct {
	Type        string `json:"type"`
	Severity    string `json:"severity"`
	TitleKey    string `json:"titleKey"`
	ImageBase64 string `json:"img"`
	OccurredAt  string `json:"occurredAt"`
	Meta        struct {
		Kind   string `json:"kind"`
		Person struct {
			FullName string `json:"fullName"`
			Gender   string `json:"gender"`
		} `json:"person"`
		RawID string `json:"rawId"`
	} `json:"meta"`
}

func (s *CentAccessPushDataAdapterImpl) PushFaceRecognitionEventData(payload *schema.FaceRecognitionEventSchema) error {

	fmt.Println("Push data to cent access !!!")
	var sendPayload = new(CentAccessFaceRecognitionEventSchema)
	sendPayload.Type = "info"
	sendPayload.Severity = "low"
	sendPayload.TitleKey = "notis.faceDetected"
	sendPayload.ImageBase64 = payload.ImageBase64
	sendPayload.OccurredAt = payload.StampDateTime.Format("2006-01-02T15:04:05Z")
	sendPayload.Meta.Kind = "face"
	sendPayload.Meta.Person.FullName = payload.PersonInformation.FirstName
	sendPayload.Meta.Person.Gender = payload.PersonInformation.Gender
	sendPayload.Meta.RawID = "face-001"

	fmt.Println(sendPayload)

	////////////////////////////
	/// Webhook 1 //////////////
	////////////////////////////

	webhookPath := "/api/debug/print"
	webhookUrl := config.DESTINATION_HOST_ADDRESS + webhookPath
	log.Println("Webhook URL:", webhookUrl)

	convertedJsonPayload, err := json.Marshal(sendPayload)
	if err != nil {
		fmt.Printf("Error marshalling JSON: %v\n", err)
		return err
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", webhookUrl, bytes.NewBuffer(convertedJsonPayload))
	if err != nil {
		fmt.Printf("Error creating HTTP request: %v\n", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Device-Key", payload.DeviceInformation.ID)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending webhook POST request: %v\n", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Webhook responded with status code: %d\n", resp.StatusCode)
	} else {
		fmt.Println("Webhook sent successfully.")
	}

	return nil
}

func (s *CentAccessPushDataAdapterImpl) PushEmergencyAlarmEventData(payload *schema.EmergencyAlarmEventSchema) error {
	log.Println("Emergency alarm event data pushed to Cent Access")
	return nil
}

func (s *CentAccessPushDataAdapterImpl) PushPM25Value(pm25Value float64, pm10Value float64, humidityValue float64, temperatureValue float64) error {
	log.Println("PM2.5 value pushed to Cent Access : " + fmt.Sprintf("%.2f", pm25Value))
	log.Println("PM10 value pushed to Cent Access : " + fmt.Sprintf("%.2f", pm10Value))
	log.Println("Humidity value pushed to Cent Access : " + fmt.Sprintf("%.2f", humidityValue))
	log.Println("Temperature value pushed to Cent Access : " + fmt.Sprintf("%.2f", temperatureValue))
	return nil
}

func (s *CentAccessPushDataAdapterImpl) PushWaterSensorData() error {
	log.Println("Water sensor data pushed to Cent Access")
	return nil
}
