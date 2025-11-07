package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/putteror/iot-gateway/internal/app/schema"
)

// Add necessary imports here

// Add service implementation here
type DahuaCameraFaceRecognitionService interface {
	FaceRecognitionEvent(paylaod *schema.ReceiveDahuaCameraFaceRecognitionPayload) error
}

type DahuaCameraFaceRecognitionServiceImpl struct {
}

func NewDahuaCameraFaceRecognitionService() DahuaCameraFaceRecognitionService {
	return &DahuaCameraFaceRecognitionServiceImpl{}
}

// Add service methods here
func (s *DahuaCameraFaceRecognitionServiceImpl) FaceRecognitionEvent(paylaod *schema.ReceiveDahuaCameraFaceRecognitionPayload) error {

	returnJsonPayload, err := json.Marshal(paylaod)
	if err != nil {
		// จัดการกับข้อผิดพลาดในการแปลง JSON
		fmt.Printf("Error marshalling JSON: %v\n", err)
		return err
	}
	// แสดงผล JSON ที่ได้
	fmt.Printf("Received Payload JSON: %s\n", string(returnJsonPayload))

	now := time.Now()
	iso8601String := now.Format(time.RFC3339)

	var sendPayload schema.SendDahuaCameraFaceRecognitionPayload
	sendPayload.Type = "info"
	sendPayload.Severity = "low"
	sendPayload.TitleKey = "notis.faceDetected"
	sendPayload.OccurredAt = iso8601String
	sendPayload.Meta.Kind = "face"
	sendPayload.Meta.Person.FullName = paylaod.Data.Name
	sendPayload.Meta.Person.Gender = "MALE"
	sendPayload.Meta.RawID = "face-001"
	err = json.Unmarshal(returnJsonPayload, &sendPayload)
	if err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err)
		return err
	}

	fmt.Printf("Send Payload: %+v\n", sendPayload)

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
