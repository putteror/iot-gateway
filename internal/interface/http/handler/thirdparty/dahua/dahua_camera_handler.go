package dahuahandler

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/putteror/iot-gateway/internal/app/schema"
	"github.com/putteror/iot-gateway/internal/app/service"
	"github.com/putteror/iot-gateway/internal/config"
	dahuaschema "github.com/putteror/iot-gateway/internal/interface/http/schema/thirdparty/dahua"
)

type DahuaCameraFaceRecognitionHandler struct {
	service service.WebhookService
}

func NewDahuaCameraFaceRecognitionHandler(service service.WebhookService) *DahuaCameraFaceRecognitionHandler {
	return &DahuaCameraFaceRecognitionHandler{service: service}
}

// payload //
type ReceiveDahuaCameraFaceRecognitionPayload struct {
	Action string `json:"Action"`
	Code   string `json:"Code"`
	Data   struct {
		Address       string `json:"Address"`
		Class         string `json:"Class"`
		Face          any    `json:"Face"`
		IsGlobalScene bool   `json:"IsGlobalScene"`
		Name          string `json:"Name"`
		Object        any    `json:"Object"`
		RealUTC       int64  `json:"RealUTC"`
		RuleID        int    `json:"RuleID"`
	} `json:"Data"`
	Index int `json:"Index"`
}

// Create creates a new access control device.
func (h *DahuaCameraFaceRecognitionHandler) FaceRecognitionEvent(c *gin.Context) {

	deviceID := c.Param("id")

	// Get body data from json request
	var bodyRequest *dahuaschema.DahuaCameraFaceRecognitionEventSchema
	if err := c.ShouldBindJSON(&bodyRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to read request body: " + err.Error(),
		})
		return
	}

	// get person data if have candidiate

	var defaultPayload = new(schema.FaceRecognitionEventSchema)
	defaultPayload.Success = true
	defaultPayload.Type = "info"
	defaultPayload.ImageBase64 = ""
	defaultPayload.StampDateTime = time.Unix(int64(bodyRequest.Events[0].Data.RealUTC), 0)

	defaultPayload.PersonInformation.Age = bodyRequest.Events[0].Data.Face.Age
	defaultPayload.PersonInformation.Gender = bodyRequest.Events[0].Data.Face.Sex

	// add and edit data if found candidate
	if len(bodyRequest.Events[0].Data.Candidates) > 0 {
		defaultPayload.Confidence = bodyRequest.Events[0].Data.Candidates[0].Similarity
		personInformation := bodyRequest.Events[0].Data.Candidates[0].Person

		if strings.Contains(personInformation.Name, " ") {
			splitName := strings.Split(personInformation.Name, " ")
			if len(splitName) == 2 {
				defaultPayload.PersonInformation.FirstName = splitName[0]
				defaultPayload.PersonInformation.LastName = splitName[1]
			} else if len(splitName) == 3 {
				defaultPayload.PersonInformation.FirstName = splitName[0]
				defaultPayload.PersonInformation.MiddleName = strings.Join(splitName[1:len(splitName)-1], " ")
				defaultPayload.PersonInformation.LastName = splitName[len(splitName)-1]
			} else {
				defaultPayload.PersonInformation.FirstName = splitName[0]
			}
		}

		defaultPayload.PersonInformation.ID = personInformation.ID
		defaultPayload.PersonInformation.Gender = personInformation.Sex
		defaultPayload.PersonInformation.Age = personInformation.Age

	}

	defaultPayload.DeviceInformation.ID = deviceID
	defaultPayload.DeviceInformation.Name = ""
	defaultPayload.DeviceInformation.IPAddress = ""
	defaultPayload.DeviceInformation.MACAddress = ""

	h.service.PushDataToDestination(defaultPayload, "face-recognition", config.DESTINATION_TYPE)

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully receive face recognition event",
	})
}

func (h *DahuaCameraFaceRecognitionHandler) FaceRecognitionImageEvent(c *gin.Context) {
	deviceID := c.Param("id")
	contentType := c.Request.Header.Get("Content-Type")

	// check multipart/x-mixed-replace header
	if !strings.HasPrefix(contentType, "multipart/x-mixed-replace") {
		log.Printf("ERROR: Invalid Content-Type: %s", contentType)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported content type. Expected multipart/x-mixed-replace"})
		return
	}

	_, params, err := mime.ParseMediaType(contentType)
	if err != nil {
		log.Printf("ERROR: Failed to parse Content-Type: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Content-Type header"})
		return
	}
	boundary := params["boundary"]
	if boundary == "" {
		log.Println("ERROR: Missing boundary in Content-Type")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing boundary in Content-Type"})
		return
	}

	var bodyReader io.Reader = c.Request.Body
	reader := multipart.NewReader(bodyReader, boundary)

	// loop check part
	faceRecognitionPayload := h.processMultipartStream(reader, deviceID)
	h.service.PushDataToDestination(faceRecognitionPayload, "face-recognition", config.DESTINATION_TYPE)

	// 5. ตอบกลับ
	c.JSON(http.StatusOK, gin.H{"message": "Face recognition stream processing initiated"})
}

func (h *DahuaCameraFaceRecognitionHandler) processMultipartStream(reader *multipart.Reader, deviceID string) (returnPayload *schema.FaceRecognitionEventSchema) {
	partIndex := 0

	var imageBase64 string
	var dahuaPayload *dahuaschema.DahuaCameraFaceRecognitionEventSchema

	for {
		// NextPart() จะบล็อกการทำงานจนกว่าจะได้รับ Part ใหม่
		part, err := reader.NextPart()

		if err == io.EOF {
			log.Println("INFO: End of multipart stream (EOF).")
			break
		}
		if err != nil {
			log.Printf("ERROR: Failed to read next part: %v", err)
			break
		}

		partIndex++
		partContentType := part.Header.Get("Content-Type")

		// JSON/TEXT Part ⭐️
		if strings.Contains(partContentType, "text/plain") || strings.Contains(partContentType, "application/json") {
			data, err := io.ReadAll(part)
			if err != nil {
				log.Printf("ERROR: Failed to read JSON part %d: %v", partIndex, err)
				return
			}
			var eventData map[string]interface{}
			if err := json.Unmarshal(data, &eventData); err == nil {
				log.Printf("INFO: Part %d (JSON Event) - Processed event data. Length: %d", partIndex, len(data))
			} else {
				log.Printf("WARNING: Part %d (JSON/Text) - Failed to unmarshal JSON: %v. Raw data length: %d", partIndex, err, len(data))
				dahuaPayload = nil
			}
			dahuaPayload = new(dahuaschema.DahuaCameraFaceRecognitionEventSchema)
			err = json.Unmarshal(data, &dahuaPayload)
			if err != nil {
				log.Printf("ERROR: Failed to unmarshal JSON part %d: %v", partIndex, err)
				return
			}

			// Image/JPEG Part ⭐️
		} else if strings.Contains(partContentType, "image/jpeg") {
			imageData, err := io.ReadAll(part)
			if err != nil {
				log.Printf("ERROR: Failed to read Image part %d: %v", partIndex, err)
				return
			}
			imageBase64 = base64.StdEncoding.EncodeToString(imageData)
			if imageBase64 != "" {
				log.Printf("Success to get image base64")
			}

		} else {
			log.Printf("WARNING: Part %d - Unknown Content-Type: %s", partIndex, partContentType)
		}
	}

	faceRecognitionPayload := h.SaveDataToDefaultFormat(dahuaPayload, imageBase64, deviceID)
	return faceRecognitionPayload
}

// function to cconvert data
func (h *DahuaCameraFaceRecognitionHandler) SaveDataToDefaultFormat(payload *dahuaschema.DahuaCameraFaceRecognitionEventSchema, imageBase64 string, deviceID string) (returnPayload *schema.FaceRecognitionEventSchema) {

	var faceRecognitionPayload = new(schema.FaceRecognitionEventSchema)

	faceRecognitionPayload.Success = true
	faceRecognitionPayload.Type = "info"
	faceRecognitionPayload.ImageBase64 = ""
	faceRecognitionPayload.StampDateTime = time.Unix(int64(payload.Events[0].Data.RealUTC), 0)

	faceRecognitionPayload.PersonInformation.Age = payload.Events[0].Data.Face.Age
	faceRecognitionPayload.PersonInformation.Gender = payload.Events[0].Data.Face.Sex

	// add and edit data if found candidate
	if len(payload.Events[0].Data.Candidates) > 0 {
		faceRecognitionPayload.Confidence = payload.Events[0].Data.Candidates[0].Similarity
		personInformation := payload.Events[0].Data.Candidates[0].Person

		if strings.Contains(personInformation.Name, " ") {
			splitName := strings.Split(personInformation.Name, " ")
			if len(splitName) == 2 {
				faceRecognitionPayload.PersonInformation.FirstName = splitName[0]
				faceRecognitionPayload.PersonInformation.LastName = splitName[1]
			} else if len(splitName) == 3 {
				faceRecognitionPayload.PersonInformation.FirstName = splitName[0]
				faceRecognitionPayload.PersonInformation.MiddleName = strings.Join(splitName[1:len(splitName)-1], " ")
				faceRecognitionPayload.PersonInformation.LastName = splitName[len(splitName)-1]
			} else {
				faceRecognitionPayload.PersonInformation.FirstName = splitName[0]
			}
		} else {
			faceRecognitionPayload.PersonInformation.FirstName = personInformation.Name
		}

		faceRecognitionPayload.PersonInformation.ID = personInformation.ID
		faceRecognitionPayload.PersonInformation.Gender = personInformation.Sex
		// faceRecognitionPayload.PersonInformation.Age = personInformation.Age // age is not match with db

	}

	faceRecognitionPayload.DeviceInformation.ID = deviceID
	faceRecognitionPayload.DeviceInformation.Name = ""
	faceRecognitionPayload.DeviceInformation.IPAddress = ""
	faceRecognitionPayload.DeviceInformation.MACAddress = ""

	// image
	faceRecognitionPayload.ImageBase64 = imageBase64

	return faceRecognitionPayload

}
