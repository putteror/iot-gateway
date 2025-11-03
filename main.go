package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/putteror/iot-gateway/http"
	"github.com/putteror/iot-gateway/http/handler"
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

	defaultHandler := handler.NewDefaultHandler()
	inboundHandler := handler.NewInboundHandler()

	appRouter := http.NewRouter(
		defaultHandler,
		inboundHandler,
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

}
