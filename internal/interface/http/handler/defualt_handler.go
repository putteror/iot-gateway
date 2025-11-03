package handler

import (
	"github.com/gin-gonic/gin"
)

type DefaultHandler struct {
}

func NewDefaultHandler() *DefaultHandler {
	return &DefaultHandler{}
}

// GetAll retrieves all access control devices.
func (h *DefaultHandler) GetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success to get api",
	})
}

// GetByID retrieves an access control device by its ID.
func (h *DefaultHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "success to get api by id :" + id,
	})
}

// Create creates a new access control device.
func (h *DefaultHandler) Create(c *gin.Context) {
	var bodyRequest interface{}
	if err := c.ShouldBindJSON(&bodyRequest); err != nil {
		c.JSON(400, gin.H{
			"message": "fail to post api : " + err.Error(),
		})
		return
	}

	// return body to json
	c.JSON(200, gin.H{
		"message": "success to post api",
		"data":    bodyRequest,
	})
}

// Update updates an existing access control device.
func (h *DefaultHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var bodyRequest interface{}
	if err := c.ShouldBindJSON(&bodyRequest); err != nil {
		c.JSON(400, gin.H{
			"message": "fail to put api : " + err.Error(),
		})
		return
	}

	// return body to json
	c.JSON(200, gin.H{
		"message": "success to put api by id :" + id,
		"data":    bodyRequest,
	})
}

// Delete deletes an access control device by its ID.
func (h *DefaultHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "success to delete api by id :" + id,
	})
}
