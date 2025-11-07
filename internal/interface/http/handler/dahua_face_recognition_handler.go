package handler

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/putteror/iot-gateway/internal/app/schema"
	"github.com/putteror/iot-gateway/internal/app/service"
)

type DahuaCameraFaceRecognitionHandler struct {
	service service.DahuaCameraFaceRecognitionService
}

func NewDahuaCameraFaceRecognitionHandler(service service.DahuaCameraFaceRecognitionService) *DahuaCameraFaceRecognitionHandler {
	return &DahuaCameraFaceRecognitionHandler{service: service}
}

// GetByID retrieves an access control device by its ID.
func (h *DahuaCameraFaceRecognitionHandler) Get(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully printed headers and parameters to console",
	})

}

// Create creates a new access control device.
func (h *DahuaCameraFaceRecognitionHandler) ReceiveFaceRecognitionEvent(c *gin.Context) {

	// Get body data from json request
	var bodyRequest *schema.ReceiveDahuaCameraFaceRecognitionPayload
	if err := c.ShouldBindJSON(&bodyRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to read request body: " + err.Error(),
		})
		return
	}

	h.service.FaceRecognitionEvent(bodyRequest)

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully printed headers, parameters, and body to console",
	})
}

func (h *DahuaCameraFaceRecognitionHandler) ReceiveFaceRecognitionImage(c *gin.Context) {
	// ... โค้ดส่วนที่ 1: ตรวจสอบ Header และ Boundary ...
	contentType := c.Request.Header.Get("Content-Type")
	_, params, err := mime.ParseMediaType(contentType)
	// ... (ตรวจสอบ error เหมือนเดิม)
	// ... (ตัดโค้ดส่วนนี้เพื่อความกระชับ)

	boundary := "--" + params["boundary"]

	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error reading request body: %v", err)})
		return
	}

	parts := bytes.Split(bodyBytes, []byte(boundary))

	// ตัวแปรสำหรับเก็บค่า EventSeq
	var eventSeq int = 0 // กำหนดค่าเริ่มต้นเป็น 0 หรือค่า Default อื่นๆ
	var imageBody []byte

	// 3. แยก JSON และ Image Part

	for _, part := range parts {
		part = bytes.TrimSpace(part)
		if len(part) == 0 || bytes.HasSuffix(part, []byte("--")) {
			continue
		}

		// 3.1 ตรวจสอบและ Unmarshal JSON Payload
		if bytes.Contains(part, []byte("Content-Type: text/plain")) || bytes.Contains(part, []byte("Content-Type: application/json")) {

			// แยก Header ออกจาก Body ของ JSON
			headerEndIndex := bytes.Index(part, []byte("\r\n\r\n"))
			if headerEndIndex == -1 {
				headerEndIndex = bytes.Index(part, []byte("\n\n"))
			}

		}

		// 3.2 ตรวจสอบและดึง Image Binary Data
		if bytes.Contains(part, []byte("Content-Type: image/jpeg")) {
			// ... (โค้ดหา headerEndIndex ของ image part เหมือนเดิม)
			// ... (โค้ดแยก imageBody ออกมาเหมือนเดิม)
			// ตัวอย่างการดึง imageBody
			headerEndIndex := bytes.Index(part, []byte("\r\n\r\n"))
			if headerEndIndex != -1 {
				imageBody = part[headerEndIndex+4:]
			} else {
				// ... การจัดการ headerEndIndex == -1 ...
			}
		}
	}

	// 4. ตรวจสอบว่ามีรูปภาพและ EventSeq ที่ต้องการหรือไม่
	if len(imageBody) == 0 {
		log.Println("No JPEG image part found in the request body.")
		c.JSON(http.StatusBadRequest, gin.H{"error": "No JPEG image part found."})
		return
	}

	if eventSeq == 0 {
		log.Println("Warning: EventSeq not found in JSON payload. Using timestamp for filename.")
	}

	// 5. บันทึกไฟล์ (ใช้ EventSeq ที่ดึงมาได้)
	uploadDir := "./uploads"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir(uploadDir, os.ModePerm)
	}

	// ใช้ EventSeq ในการตั้งชื่อไฟล์ หรือใช้ Unix Timestamp หาก EventSeq เป็น 0
	var fileName string
	if eventSeq != 0 {
		fileName = "faceImage.jpg"
	} else {
		fileName = "faceImage.jpg" // หรือใช้ time.Now().Unix()
	}

	savePath := filepath.Join(uploadDir, fileName)

	if err := os.WriteFile(savePath, imageBody, 0644); err != nil {
		log.Printf("Failed to save image file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to save file: %v", err)})
		return
	}

	fmt.Printf("✅ File saved successfully to %s (Size: %d bytes)\n", savePath, len(imageBody))

	c.JSON(http.StatusOK, gin.H{
		"message":  "Image received and saved successfully",
		"saved_to": savePath,
		"size":     len(imageBody),
	})
}
