package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/putteror/iot-gateway/internal/adapter"
	"github.com/putteror/iot-gateway/internal/adapter/push"
	"github.com/putteror/iot-gateway/internal/app/service"
	"github.com/putteror/iot-gateway/internal/interface/http"
	"github.com/putteror/iot-gateway/internal/interface/http/handler"
	dahuahandler "github.com/putteror/iot-gateway/internal/interface/http/handler/thirdparty/dahua"
	hikvisionhandler "github.com/putteror/iot-gateway/internal/interface/http/handler/thirdparty/hikvision"
)

func main() {

	pm25Fetcher := adapter.NewGetPM25Impl()
	zytaAdapter := push.NewPushDataServiceImpl()
	centAccessAdapter := push.NewCentAccessPushDataServiceImpl()

	retentionService := service.NewRetentionService(pm25Fetcher, zytaAdapter, centAccessAdapter)
	if err := retentionService.RetentionGetPM25AndPushToDestination(); err != nil {
		log.Fatalf("Failed to start PM2.5 background job: %v", err)
	}
	if err := retentionService.RetentionGetWaterSensorAndPushToDestination(); err != nil {
		log.Fatalf("Failed to start Water Sensor background job: %v", err)
	}

	err := godotenv.Load()
	if err != nil {
		log.Println("Note: Error loading .env file. Proceeding with system environment variables.")
	}

	apiServicePort := os.Getenv("API_SERVICE_PORT")
	if apiServicePort == "" {
		apiServicePort = "8080"
	}

	webhookService := service.NewWebhookService(
		push.NewPushDataServiceImpl(),
		push.NewCentAccessPushDataServiceImpl(),
	)

	defaultHandler := handler.NewDefaultHandler()
	thirdPartyDahuaCameraFaceRecognitionHandler := dahuahandler.NewDahuaCameraFaceRecognitionHandler(webhookService)
	thirdPartyHikvisionCameraEmergencyAlarmHandler := hikvisionhandler.NewHikvisionCameraEmergencyAlarmHandler(webhookService)

	appRouter := http.NewRouter(
		thirdPartyDahuaCameraFaceRecognitionHandler,
		thirdPartyHikvisionCameraEmergencyAlarmHandler,
		defaultHandler,
	)

	log.Printf("Server is starting on port %s", apiServicePort)
	fmt.Println("==========================================================")
	fmt.Println("==== If you want to test receiving API from 3rd-party")
	fmt.Println("====        set this url to 3rd-party device")
	fmt.Printf("====   http://< this server ip >:%s/api/inbound\n", apiServicePort)
	fmt.Println("==========================================================")
	if err := appRouter.Run(":" + apiServicePort); err != nil {
		log.Fatalf("Could not listen on %s: %v\n", apiServicePort, err)
	}

	// For testing purpose
	// integration.TestGetData()

}
