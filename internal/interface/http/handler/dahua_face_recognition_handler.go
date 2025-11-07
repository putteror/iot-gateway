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
	"strings"

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
	contentType := c.Request.Header.Get("Content-Type")
	_, params, err := mime.ParseMediaType(contentType)
	if err != nil || params["boundary"] == "" || !strings.Contains(contentType, "multipart/x-mixed-replace") {
		log.Printf("Invalid Content-Type or missing boundary: %s", contentType)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Content-Type. Expected multipart/x-mixed-replace with boundary."})
		return
	}

	// สร้างสตริง Boundary ที่ใช้จริง (เช่น --myboundary)
	boundary := "--" + params["myboundary"]

	// 2. อ่าน Request Body ทั้งหมด
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error reading request body: %v", err)})
		return
	}

	// เนื่องจากข้อมูลไบนารีอาจเสียหายหากแปลงเป็น string เราจึงทำงานบน []byte โดยตรง
	// เราจะใช้ bytes.Split เพื่อแบ่งข้อมูลตาม Boundary

	// 3. แบ่ง Body ออกเป็นส่วนๆ ตาม Boundary
	// Parts[0] จะเป็น preamble ก่อน Boundary แรก
	// Parts[1] จะเป็น JSON Payload
	// Parts[2] จะเป็น Image Payload
	parts := bytes.Split(bodyBytes, []byte(boundary))

	// 4. วนหาและประมวลผลส่วนที่เป็น Image/JPEG
	for _, part := range parts {
		// Trim carriage returns/newlines ที่อาจมี
		part = bytes.TrimSpace(part)
		if len(part) == 0 || bytes.HasSuffix(part, []byte("--")) {
			continue // ข้ามส่วนว่างเปล่าหรือส่วนท้าย (EOF)
		}

		// ตรวจสอบว่าเป็น Image/JPEG หรือไม่
		if bytes.Contains(part, []byte("Content-Type: image/jpeg")) {
			// 4.1 แยก Header และ Body
			// หาจุดสิ้นสุดของ Header (ดับเบิ้ล Newline/Carriage Return: \r\n\r\n)
			// ใช้ \n\n หรือ \r\n\r\n ในการหา ขึ้นอยู่กับ client
			headerEndIndex := bytes.Index(part, []byte("\r\n\r\n"))
			if headerEndIndex == -1 {
				// ลองใช้ \n\n (อาจจะเกิดจากระบบปฏิบัติการที่ใช้)
				headerEndIndex = bytes.Index(part, []byte("\n\n"))
				if headerEndIndex == -1 {
					log.Println("Could not find end of header in image part.")
					continue
				}
			}

			// 4.2 ดึง Binary Data ของรูปภาพ
			// ข้อมูลรูปภาพจะอยู่หลัง \r\n\r\n (ยาว 4 ไบต์) หรือ \n\n (ยาว 2 ไบต์)
			var imageBody []byte
			if bytes.HasPrefix(part[headerEndIndex:], []byte("\r\n\r\n")) {
				imageBody = part[headerEndIndex+4:]
			} else if bytes.HasPrefix(part[headerEndIndex:], []byte("\n\n")) {
				imageBody = part[headerEndIndex+2:]
			} else {
				imageBody = part[headerEndIndex:] // กรณีที่ body ต่อทันที
			}

			// 5. บันทึกไฟล์
			// ตรวจสอบหรือสร้าง directory
			uploadDir := "./uploads"
			if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
				if err := os.Mkdir(uploadDir, os.ModePerm); err != nil {
					log.Printf("Failed to create directory %s: %v", uploadDir, err)
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
					return
				}
			}

			// สร้างชื่อไฟล์ที่ไม่ซ้ำกัน (ใช้ RealUTC หรือ EventID จาก JSON มาช่วยได้)
			fileName := fmt.Sprintf("faceImage_%d.jpeg", c.MustGet("EventSeq").(int)) // สมมติว่าดึง EventSeq จาก JSON มาเก็บใน Context แล้ว
			savePath := filepath.Join(uploadDir, fileName)

			if err := os.WriteFile(savePath, imageBody, 0644); err != nil {
				log.Printf("Failed to save image file: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to save file: %v", err)})
				return
			}

			fmt.Printf("✅ File saved successfully to %s (Size: %d bytes)\n", savePath, len(imageBody))

			// 6. ส่ง response กลับและหยุดการทำงาน (สมมติว่าต้องการแค่รูปเดียว)
			c.JSON(http.StatusOK, gin.H{
				"message":  "Image received and saved successfully from multipart/x-mixed-replace",
				"saved_to": savePath,
				"size":     len(imageBody),
			})
			return
		}
	}

	// ถ้าไม่พบรูปภาพ JPEG เลย
	log.Println("No JPEG image part found in the request body.")
	c.JSON(http.StatusBadRequest, gin.H{"error": "No JPEG image part found in the request body."})
}
