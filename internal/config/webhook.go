package config

import (
	"os"

	"github.com/joho/godotenv"
)

var WEBHOOK_URL string

func LoadWebhookConfig() {
	godotenv.Load()
	WEBHOOK_URL = os.Getenv("WEBHOOK_URL")
}
