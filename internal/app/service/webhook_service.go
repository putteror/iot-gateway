package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/putteror/iot-gateway/internal/config"
)

type WebhookService interface {
	WebhookByPassEvent(paylaod interface{}) string
}

type WebhookServiceImpl struct {
}

func NewWebhookService() WebhookService {
	return &WebhookServiceImpl{}
}

func (s *WebhookServiceImpl) WebhookByPassEvent(paylaod interface{}) string {
	fmt.Println(paylaod)

	jsonPayload, err := json.Marshal(paylaod)
	if err != nil {
		// จัดการกับข้อผิดพลาดในการแปลง JSON
		fmt.Printf("Error marshalling JSON: %v\n", err)
		return "Error: Failed to marshal JSON"
	}

	println(config.WEBHOOK_URL)

	// 2. ส่ง HTTP POST request
	resp, err := http.Post(config.WEBHOOK_URL, "application/json", bytes.NewBuffer(jsonPayload))
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
