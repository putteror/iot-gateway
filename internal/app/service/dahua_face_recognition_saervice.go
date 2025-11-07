package service

import (
	"encoding/json"
	"fmt"

	"github.com/putteror/iot-gateway/internal/app/schema"
)

// Add necessary imports here

// Add service implementation here
type DahuaCameraFaceRecognitionService interface {
	FaceRecognitionEvent(paylaod *schema.DahuaNVRFaceRecognitionEvent) error
}

type DahuaCameraFaceRecognitionServiceImpl struct {
}

func NewDahuaCameraFaceRecognitionService() DahuaCameraFaceRecognitionService {
	return &DahuaCameraFaceRecognitionServiceImpl{}
}

// Add service methods here
func (s *DahuaCameraFaceRecognitionServiceImpl) FaceRecognitionEvent(paylaod *schema.DahuaNVRFaceRecognitionEvent) error {

	jsonPayload, err := json.Marshal(paylaod)
	if err != nil {
		// จัดการกับข้อผิดพลาดในการแปลง JSON
		fmt.Printf("Error marshalling JSON: %v\n", err)
		return err
	}

	//print jsonPayload to console
	fmt.Println("JSON Payload:", string(jsonPayload))

	/*

		webhookUrl := config.WEBHOOK_HOST_ADDRESS + config.WEBHOOK_FACE_RECOGNITION_PATH
		log.Println("Webhook URL:", webhookUrl)

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
	*/

	return nil
}
