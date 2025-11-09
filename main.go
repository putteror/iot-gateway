package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/putteror/iot-gateway/internal/adapter/push"
	"github.com/putteror/iot-gateway/internal/app/service"
	"github.com/putteror/iot-gateway/internal/interface/http"
	"github.com/putteror/iot-gateway/internal/interface/http/handler"
	dahuahandler "github.com/putteror/iot-gateway/internal/interface/http/handler/thirdparty/dahua"
)

func main() {

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

	appRouter := http.NewRouter(
		thirdPartyDahuaCameraFaceRecognitionHandler,
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
