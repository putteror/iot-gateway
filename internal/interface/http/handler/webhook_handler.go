package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/putteror/iot-gateway/internal/app/service"
)

type WebhookHandler struct {
	service service.WebhookService
}

func NewWebhookHandler(service service.WebhookService) *WebhookHandler {
	return &WebhookHandler{service: service}
}

// Add handler methods here
func (h *WebhookHandler) ByPassPost(c *gin.Context) {
	// Get body data from json request
	var bodyRequest interface{}
	if err := c.ShouldBindJSON(&bodyRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to read request body: " + err.Error(),
		})
		return
	}

	h.service.WebhookByPassEvent(bodyRequest)

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully printed headers, parameters, and body to console",
	})
}
