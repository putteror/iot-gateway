package hikvisionhandler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/putteror/iot-gateway/internal/app/schema"
	"github.com/putteror/iot-gateway/internal/app/service"
	"github.com/putteror/iot-gateway/internal/config"
	hikvision "github.com/putteror/iot-gateway/internal/interface/http/schema/thirdparty/hikvision"
)

type HikvisionCameraEmergencyAlarmHandler struct {
	service service.WebhookService
}

func NewHikvisionCameraEmergencyAlarmHandler(service service.WebhookService) *HikvisionCameraEmergencyAlarmHandler {
	return &HikvisionCameraEmergencyAlarmHandler{service: service}
}

// Create creates a new access control device.
func (h *HikvisionCameraEmergencyAlarmHandler) EmergencyAlarmEvent(c *gin.Context) {

	deviceID := c.Param("id")
	siteID := c.Param("siteId")

	// Get body data from json request
	var bodyRequest *hikvision.HikvisionCameraEmergencyAlarmEventSchema
	if err := c.ShouldBindJSON(&bodyRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to read request body: " + err.Error(),
		})
		return
	}

	var defaultPayload = new(schema.EmergencyAlarmEventSchema)
	defaultPayload.Type = "info"
	defaultPayload.StampDateTime = time.Unix(int64(bodyRequest.Data.UTC), 0)
	defaultPayload.SiteID = siteID
	defaultPayload.DeviceInformation.ID = deviceID
	h.service.PushDataToDestination(defaultPayload, "emergency-alarm", config.DESTINATION_TYPE)

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully receive face recognition event",
	})
}
