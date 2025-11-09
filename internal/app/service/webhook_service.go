package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/putteror/iot-gateway/internal/adapter/push"
	"github.com/putteror/iot-gateway/internal/app/schema"
	"github.com/putteror/iot-gateway/internal/config"
)

type WebhookService interface {
	PushDataToDestination(paylaod interface{}, eventType string, ThirdPartyName string) string
	WebhookByPassEvent(paylaod interface{}) string
}

// set if WEBHOOK_SERVICE_NAME = "zyta"

type WebhookServiceImpl struct {
	DestinationAdapters map[string]push.PushDataAdapter
}

func NewWebhookService(
	zytaAdapter push.ZytaPushDataAdapter,
	centAccessAdapter push.CentAccessPushDataAdapter,
) WebhookService {
	return &WebhookServiceImpl{
		DestinationAdapters: map[string]push.PushDataAdapter{
			"zyta":        zytaAdapter,       // Key คือ "zyta"
			"cent-access": centAccessAdapter, // Key คือ "cent-access"
		},
	}
}

func (s *WebhookServiceImpl) PushDataToDestination(payload interface{}, eventType string, DestinationType string) string {

	adapter, ok := s.DestinationAdapters[DestinationType]
	if !ok {
		message := fmt.Sprintf("Error: Third party adapter '%s' not supported.", DestinationType)
		fmt.Println(message)
		return message
	}

	// 2. เรียกใช้เมธอดของ Adapter ที่ดึงมา
	switch eventType {
	case "face-recognition":
		adapter.PushFaceRecognitionEventData(payload.(*schema.FaceRecognitionEventSchema))
	default:
		return "Processing complete"
	}

	return "Processing complete"
}

func (s *WebhookServiceImpl) WebhookByPassEvent(paylaod interface{}) string {
	fmt.Println(paylaod)

	jsonPayload, err := json.Marshal(paylaod)
	if err != nil {
		// จัดการกับข้อผิดพลาดในการแปลง JSON
		fmt.Printf("Error marshalling JSON: %v\n", err)
		return "Error: Failed to marshal JSON"
	}

	webhookUrl := config.DESTINATION_HOST_ADDRESS

	// 2. ส่ง HTTP POST request
	resp, err := http.Post(webhookUrl, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Printf("Error sending webhook POST request: %v\n", err)
		return "Error: Failed to send webhook"
	}
	// 3. ปิด body ของ response เพื่อไม่ให้เกิด resource leak
	defer resp.Body.Close()

	// 4. ตรวจสอบสถานะของ HTTP response (ไม่จำเป็นต้องอ่าน body ถ้าไม่ต้องการ)
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Webhook responded with status code: %d\n", resp.StatusCode)
		// คุณอาจต้องการอ่าน response body เพื่อดูข้อความผิดพลาดเพิ่มเติม
	} else {
		fmt.Println("Webhook sent successfully.")
	}

	return "Face Recognition Event from Dahua NVR"
}
