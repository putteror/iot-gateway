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

	// returnJsonPayload, err := json.Marshal(paylaod)
	// if err != nil {
	// 	// จัดการกับข้อผิดพลาดในการแปลง JSON
	// 	fmt.Printf("Error marshalling JSON: %v\n", err)
	// 	return err
	// }
	// แสดงผล JSON ที่ได้
	// fmt.Printf("Received Payload JSON: %s\n", string(returnJsonPayload))

	// convert image form /uploads/faceImage.jpg to base64 string
	filePath := "./uploads/faceImage.jpg"
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filePath, err)
		return err
	}
	base64String := base64.StdEncoding.EncodeToString(fileData)

	// Prepare the webhook URL
	now := time.Now()
	iso8601String := now.Format(time.RFC3339)

	var sendFaceRecPayload schema.SendDahuaCameraFaceRecognitionPayload
	sendFaceRecPayload.Type = "info"
	sendFaceRecPayload.Severity = "low"
	sendFaceRecPayload.TitleKey = "notis.faceDetected"
	sendFaceRecPayload.ImageBase64 = base64String
	sendFaceRecPayload.OccurredAt = iso8601String
	sendFaceRecPayload.Meta.Kind = "face"
	sendFaceRecPayload.Meta.Person.FullName = paylaod.Data.Name
	sendFaceRecPayload.Meta.Person.Gender = "MALE"
	sendFaceRecPayload.Meta.RawID = "face-001"

	var sendMotionPayload schema.SendDuhaCameraMotionDetectPayload
	sendMotionPayload.Type = "info"
	sendMotionPayload.Severity = "low"
	sendMotionPayload.TitleKey = "notis.motionDetected"
	sendMotionPayload.ImageBase64 = base64String
	sendMotionPayload.OccurredAt = iso8601String
	sendMotionPayload.Meta.Category = "fall"
	sendMotionPayload.Meta.TrackedBy = "smartpole"
	sendMotionPayload.Meta.Confidence = 0.94

	// แสดงผล JSON ที่ได้
	// fmt.Printf("Converted Payload JSON: %s\n", string(convertedJsonPayload))

	////////////////////////////
	/// Webhook 1 //////////////
	////////////////////////////

	webhookUrl := config.WEBHOOK_HOST_ADDRESS + config.WEBHOOK_PATH
	log.Println("Webhook URL:", webhookUrl)

	convertedJsonPayload, err := json.Marshal(sendFaceRecPayload)
	if err != nil {
		// จัดการกับข้อผิดพลาดในการแปลง JSON
		fmt.Printf("Error marshalling JSON: %v\n", err)
		return err
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", webhookUrl, bytes.NewBuffer(convertedJsonPayload))
	if err != nil {
		fmt.Printf("Error creating HTTP request: %v\n", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Device-Key", "FACE-REC-F01")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending webhook POST request: %v\n", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Webhook responded with status code: %d\n", resp.StatusCode)
	} else {
		fmt.Println("Webhook sent successfully.")
	}

	////////////////////////////
	/// Webhook 2 //////////////
	////////////////////////////

	motioJsonPayload, err := json.Marshal(sendMotionPayload)
	if err != nil {
		// จัดการกับข้อผิดพลาดในการแปลง JSON
		fmt.Printf("Error marshalling JSON: %v\n", err)
		return err
	}

	webhook2Url := config.WEBHOOK_HOST_ADDRESS + config.WEBHOOK_PATH
	log.Println("Webhook URL:", webhookUrl)

	req, err = http.NewRequest("POST", webhook2Url, bytes.NewBuffer(motioJsonPayload))
	if err != nil {
		fmt.Printf("Error creating HTTP request: %v\n", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Device-Key", "FACE-REC-F01")

	resp, err = client.Do(req)
	if err != nil {
		fmt.Printf("Error sending webhook POST request: %v\n", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Webhook responded with status code: %d\n", resp.StatusCode)
	} else {
		fmt.Println("Webhook sent successfully.")
	}

	return nil
}
