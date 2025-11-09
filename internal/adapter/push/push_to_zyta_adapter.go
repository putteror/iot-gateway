package push

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/putteror/iot-gateway/internal/app/schema"
	"github.com/putteror/iot-gateway/internal/config"
)

type ZytaPushDataAdapter interface {
	PushFaceRecognitionEventData(payload *schema.FaceRecognitionEventSchema) error
}

type ZytaPushDataServiceImpl struct {
}

func NewPushDataServiceImpl() ZytaPushDataAdapter {
	return &ZytaPushDataServiceImpl{}
}

type ZytaFaceRecognitionEventSchema struct {
	Type        string `json:"type"`
	Severity    string `json:"severity"`
	TitleKey    string `json:"titleKey"`
	ImageBase64 string `json:"img"`
	OccurredAt  string `json:"occurredAt"`
	Meta        struct {
		Kind   string `json:"kind"`
		Person struct {
			FullName string `json:"fullName"`
			Gender   string `json:"gender"`
			Age      int    `json:"age"`
		} `json:"person"`
		RawID string `json:"rawId"`
	} `json:"meta"`
}

func (s *ZytaPushDataServiceImpl) PushFaceRecognitionEventData(payload *schema.FaceRecognitionEventSchema) error {

	var sendPayload = new(ZytaFaceRecognitionEventSchema)
	sendPayload.Type = "info"
	sendPayload.Severity = "low"
	sendPayload.TitleKey = "notis.faceDetected"
	sendPayload.ImageBase64 = payload.ImageBase64
	sendPayload.OccurredAt = payload.StampDateTime.Format("2006-01-02T15:04:05Z")
	sendPayload.Meta.Kind = "face"
	sendPayload.Meta.Person.FullName = payload.PersonInformation.FirstName + " " + payload.PersonInformation.LastName
	sendPayload.Meta.Person.Age = payload.PersonInformation.Age
	sendPayload.Meta.Person.Gender = payload.PersonInformation.Gender
	sendPayload.Meta.RawID = payload.PersonInformation.ID

	////////////////////////////
	/// Webhook 1 //////////////
	////////////////////////////

	webhookPath := "/api/webhooks/notis"
	webhookUrl := config.DESTINATION_HOST_ADDRESS + webhookPath
	log.Println("Push data to URL:", webhookUrl)

	convertedJsonPayload, err := json.Marshal(sendPayload)
	if err != nil {
		fmt.Printf("Error marshalling JSON: %v\n", err)
		return err
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", webhookUrl, bytes.NewBuffer(convertedJsonPayload))
	if err != nil {
		log.Printf("Error creating HTTP request: %v\n", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Device-Key", payload.DeviceInformation.ID)

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending webhook POST request: %v\n", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Webhook responded with status code: %d\n", resp.StatusCode)
	} else {
		log.Println("Webhook sent successfully.")
	}

	return nil
}
