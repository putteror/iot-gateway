package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/putteror/iot-gateway/internal/app/service"
)

type DebugHandler struct {
	service service.WebhookService
}

func NewDefaultHandler() *DebugHandler {
	return &DebugHandler{}
}

func (h *DebugHandler) PushRawData(c *gin.Context) {
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

func (h *DebugHandler) ConsolePrint(c *gin.Context) {

	fmt.Println("\n=========== Start Print Request =============")

	// print method
	fmt.Println("\n--- Request Method ---")
	fmt.Println(c.Request.Method)

	// print path

	fmt.Println("\n--- Request Headers --- ")
	for key, values := range c.Request.Header {
		fmt.Printf("%s: %s\n", key, values)
	}

	fmt.Println("\n--- Query Parameters (URL) ---")
	queryMap := c.Request.URL.Query()
	for key, values := range queryMap {
		fmt.Printf("%s: %s\n", key, values)
	}

	fmt.Println("\n--- ID ---")
	id := c.Param("id")
	if id != "" {
		fmt.Println(id)
	}

	bodyBytes, err := c.GetRawData()

	fmt.Println("\n--- Raw Request Body ---")

	if err != nil {
		fmt.Printf("Error reading raw body: %s\n", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to read request body",
			"error":   err.Error(),
		})
		return
	}

	// convert byte to string for return
	rawBody := string(bodyBytes)
	fmt.Println(rawBody)

	fmt.Println("\n=========== End Print Request =============")

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully printed headers and parameters to console",
	})
}
