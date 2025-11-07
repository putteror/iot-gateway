package handler

import (
	"fmt"
	"log"
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
	file, err := c.FormFile("image") // 'image' คือชื่อ field name ของไฟล์ใน form-data
	if err != nil {
		log.Println("Failed to get file 'image' . Ensure the file field name is correct.")
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Failed to get file 'image' from form: %v. Ensure the file field name is correct.", err)})
		return
	}

	// 2. รับข้อมูล Metadata (ถ้ามี)
	// ถ้ามีการส่งข้อมูล JSON/Text อื่นๆ มาใน Multipart form field อื่นๆ
	// ตัวอย่าง: ดึงค่าจาก form field ชื่อ "metadata"
	// metadata := c.PostForm("metadata")
	// fmt.Printf("Received Metadata (if any): %s\n", metadata)

	// 3. ประมวลผลไฟล์ที่ได้รับ

	// สร้างชื่อไฟล์ที่ไม่ซ้ำกันเพื่อบันทึก (ตัวอย่าง)
	fileName := "faceImage" + file.Filename
	savePath := filepath.Join("./uploads", fileName) // กำหนด path ที่จะบันทึก

	// ตรวจสอบหรือสร้าง directory
	if _, err := os.Stat("./uploads"); os.IsNotExist(err) {
		os.Mkdir("./uploads", os.ModePerm)
	}

	// บันทึกไฟล์ที่อัปโหลด
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to save file: %v", err)})
		return
	}

	fmt.Printf("File %s uploaded successfully to %s\n", file.Filename, savePath)

	// 4. ส่ง response กลับ
	c.JSON(http.StatusOK, gin.H{
		"message":  "Image received and saved successfully",
		"filename": file.Filename,
		"size":     file.Size,
		"saved_to": savePath,
	})
}
