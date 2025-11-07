package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	var bodyRequest interface{}
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
