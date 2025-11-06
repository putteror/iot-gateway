package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/putteror/iot-gateway/internal/app/schema"
)

// Add necessary imports here

// Add service implementation here
type HikvisionEmergencyAlarmService interface {
	EmergencyAlarmService(paylaod *schema.HikvisionEmergencyAlarmEvent) error
}

type HikvisionEmergencyAlarmServiceImpl struct {
}

func NewHikvisionEmergencyAlarmService() HikvisionEmergencyAlarmService {
	return &HikvisionEmergencyAlarmServiceImpl{}
}

// Add service methods here
func (s *HikvisionEmergencyAlarmServiceImpl) EmergencyAlarmService(paylaod *schema.HikvisionEmergencyAlarmEvent) error {

	jsonPayload, err := json.Marshal(paylaod)
	if err != nil {
		// จัดการกับข้อผิดพลาดในการแปลง JSON
		fmt.Printf("Error marshalling JSON: %v\n", err)
		return err
	}

	webhookUrl := "http://example.com/hikvision/emergency" // เปลี่ยนเป็น URL ของ webhook ที่ต้องการส่งข้อมูลไป
	// log.Println("Webhook URL:", webhookUrl)

	// 2. ส่ง HTTP POST request
	resp, err := http.Post(webhookUrl, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Printf("Error sending webhook POST request: %v\n", err)
		return err
	}
	// 3. ปิด body ของ response เพื่อไม่ให้เกิด resource leak
	defer resp.Body.Close()

	// 4. ตรวจสอบสถานะของ HTTP response (ไม่จำเป็นต้องอ่าน body ถ้าไม่ต้องการ)
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Webhook responded with status code: %d\n", resp.StatusCode)
	} else {
		fmt.Println("Webhook sent successfully.")
	}

	return nil
}
