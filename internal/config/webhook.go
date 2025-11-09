package config

import (
	"os"

	"github.com/joho/godotenv"
)

var DESTINATION_TYPE string
var DESTINATION_HOST_ADDRESS string

func LoadWebhookConfig() {
	godotenv.Load()
	DESTINATION_TYPE = os.Getenv("DESTINATION_TYPE")
	DESTINATION_HOST_ADDRESS = os.Getenv("DESTINATION_HOST_ADDRESS")
}
