package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/putteror/iot-gateway/internal/app/schema"
	"github.com/putteror/iot-gateway/internal/app/service"
)

type DahuaNVRHandler struct {
	service service.DahuaNVRService
}

func NewDahuaNVRHandler(service service.DahuaNVRService) *DahuaNVRHandler {
	return &DahuaNVRHandler{service: service}
}

// GetByID retrieves an access control device by its ID.
func (h *DahuaNVRHandler) Get(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully printed headers and parameters to console",
	})

}

// Create creates a new access control device.
func (h *DahuaNVRHandler) ReceiveFaceRecognitionEvent(c *gin.Context) {

	// Get body data from json request
	var bodyRequest *schema.DahuaNVRFaceRecognitionEvent
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

func (h *DahuaNVRHandler) ReceiveLicensePlateRecognitionEvent(c *gin.Context) {
	// Get body data from json request
	var bodyRequest *schema.DahuaNVRLicensePlateRecognitionEvent
	if err := c.ShouldBindJSON(&bodyRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to read request body: " + err.Error(),
		})
		return
	}

	h.service.LicensePlateRecognitionEvent(bodyRequest)

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully printed headers, parameters, and body to console",
	})
}
