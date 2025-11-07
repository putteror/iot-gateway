package config

import (
	"os"

	"github.com/joho/godotenv"
)

var WEBHOOK_HOST_ADDRESS string
var WEBHOOK_FACE_RECOGNITION_PATH string
var WEBHOOK_LICENSE_PLATE_RECOGNITION_PATH string
var WEBHOOK_PATH string

func LoadWebhookConfig() {
	godotenv.Load()
	WEBHOOK_HOST_ADDRESS = os.Getenv("WEBHOOK_HOST_ADDRESS")
	WEBHOOK_PATH = os.Getenv("WEBHOOK_PATH")
}
