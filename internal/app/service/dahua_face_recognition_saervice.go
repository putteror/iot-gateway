package service

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/putteror/iot-gateway/internal/app/schema"
	"github.com/putteror/iot-gateway/internal/config"
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
	// convert image form /uploads/faceImage.jpg to base64 string
	filePath := "./uploads/FaceImage.png"
	// 1. อ่านไฟล์เป็น []byte
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		// จัดการกับข้อผิดพลาดในการอ่านไฟล์
		fmt.Printf("Error reading file %s: %v\n", filePath, err)
		// คุณอาจเลือก return err หรือทำอย่างอื่น ขึ้นอยู่กับการจัดการข้อผิดพลาดที่คุณต้องการ
		return err
	}
	// 2. แปลง []byte ที่อ่านได้ให้เป็น Base64 string
	base64String := base64.StdEncoding.EncodeToString(fileData)

	// 3. แสดงผล Base64 string (หรือนำไปใช้ตามความต้องการ)
	fmt.Printf("Base64 String: %s\n", base64String)

	// Prepare the webhook URL
	now := time.Now()
	iso8601String := now.Format(time.RFC3339)

	var sendPayload schema.SendDahuaCameraFaceRecognitionPayload
	sendPayload.Type = "info"
	sendPayload.Severity = "low"
	sendPayload.TitleKey = "notis.faceDetected"
	sendPayload.ImageBase64 = base64String
	sendPayload.OccurredAt = iso8601String
	sendPayload.Meta.Kind = "face"
	sendPayload.Meta.Person.FullName = paylaod.Data.Name
	sendPayload.Meta.Person.Gender = "MALE"
	sendPayload.Meta.RawID = "face-001"

	convertedJsonPayload, err := json.Marshal(sendPayload)
	if err != nil {
		// จัดการกับข้อผิดพลาดในการแปลง JSON
		fmt.Printf("Error marshalling JSON: %v\n", err)
		return err
	}
	// แสดงผล JSON ที่ได้
	fmt.Printf("Converted Payload JSON: %s\n", string(convertedJsonPayload))

	webhookUrl := config.WEBHOOK_HOST_ADDRESS + config.WEBHOOK_PATH
	log.Println("Webhook URL:", webhookUrl)

	// add X-Device-Key to header
	// สร้าง HTTP client
	client := &http.Client{}

	// 1. สร้าง HTTP request
	req, err := http.NewRequest("POST", webhookUrl, bytes.NewBuffer(convertedJsonPayload))
	if err != nil {
		fmt.Printf("Error creating HTTP request: %v\n", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Device-Key", "FACE-REC-F01")

	// 2. ส่ง HTTP request
	resp, err := client.Do(req)
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
