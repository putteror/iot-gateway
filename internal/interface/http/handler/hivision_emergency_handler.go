package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HikvisionEmergencyHandler struct {
}

func NewHikvisionEmergencyHandler() *HikvisionEmergencyHandler {
	return &HikvisionEmergencyHandler{}
}

func (h *HikvisionEmergencyHandler) Get(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("\n--- ID ---")
	fmt.Println(id)

	fmt.Println("\n--- Request Headers ---")
	for key, values := range c.Request.Header {
		fmt.Printf("%s: %s\n", key, values)
	}

	fmt.Println("\n--- Query Parameters (URL) ---")
	queryMap := c.Request.URL.Query()
	for key, values := range queryMap {
		fmt.Printf("%s: %s\n", key, values)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully printed headers and parameters to console",
	})

}

// Create creates a new access control device.
func (h *HikvisionEmergencyHandler) Post(c *gin.Context) {
	fmt.Println("\n--- Request Headers --- ")
	for key, values := range c.Request.Header {
		fmt.Printf("%s: %s\n", key, values)
	}

	fmt.Println("\n--- Query Parameters (URL) ---")
	queryMap := c.Request.URL.Query()
	for key, values := range queryMap {
		fmt.Printf("%s: %s\n", key, values)
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

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully printed headers and parameters to console",
	})
}

// Update updates an existing access control device.
func (h *HikvisionEmergencyHandler) Put(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("\n--- ID ---")
	fmt.Println(id)

	fmt.Println("\n--- Request Headers --- ")
	for key, values := range c.Request.Header {
		fmt.Printf("%s: %s\n", key, values)
	}

	fmt.Println("\n--- Query Parameters (URL) ---")
	queryMap := c.Request.URL.Query()
	for key, values := range queryMap {
		fmt.Printf("%s: %s\n", key, values)
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

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully printed headers and parameters to console",
	})
}

// Delete deletes an access control device by its ID.
func (h *HikvisionEmergencyHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("\n--- ID ---")
	fmt.Println(id)

	fmt.Println("\n--- Request Headers --- ")
	for key, values := range c.Request.Header {
		fmt.Printf("%s: %s\n", key, values)
	}

	fmt.Println("\n--- Query Parameters (URL) ---")
	queryMap := c.Request.URL.Query()
	for key, values := range queryMap {
		fmt.Printf("%s: %s\n", key, values)
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

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully printed headers and parameters to console",
	})
}
