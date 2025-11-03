package config

import (
	"log"
)

func init() {
	LoadWebhookConfig()
	log.Println("Initializing config... Webhook URL loaded.")
}
