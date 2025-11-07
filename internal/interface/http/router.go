package http

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/putteror/iot-gateway/internal/interface/http/handler"
)

func NewRouter(
	defaultHandler *handler.DefaultHandler,
	hikvisionEmergencyHandler *handler.HikvisionEmergencyHandler,
	InboundHandler *handler.InboundHandler,
	dahuaNVRHandler *handler.DahuaNVRHandler,
	dahuaCameraFaceRecognitionHandler *handler.DahuaCameraFaceRecognitionHandler,
	webhookHandler *handler.WebhookHandler,
) *gin.Engine {
	router := gin.Default()

	config := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	// Apply the middleware
	router.Use(cors.New(config))
	api := router.Group("/api")
	{

		// Access Control Device endpoints
		Default := api.Group("/default")
		{
			Default.GET("/", defaultHandler.GetAll)
			Default.GET("/:id", defaultHandler.GetByID)
			Default.POST("/", defaultHandler.Create)
			Default.PUT("/:id", defaultHandler.Update)
			Default.PATCH("/", defaultHandler.Update)
			Default.DELETE("/:id", defaultHandler.Delete)
		}

		hikvisionEmergency := api.Group("/hikvision")
		{
			hikvisionEmergency.POST("/emergency", hikvisionEmergencyHandler.ReceiveEmergencyAlarmEvent)
			hikvisionEmergency.GET("/", hikvisionEmergencyHandler.Get)
			hikvisionEmergency.POST("/", hikvisionEmergencyHandler.Post)
			hikvisionEmergency.PATCH("/", hikvisionEmergencyHandler.Put)
			hikvisionEmergency.DELETE("/:id", hikvisionEmergencyHandler.Delete)
		}

		Inbound := api.Group("/inbound")
		{
			Inbound.GET("/", InboundHandler.Get)
			Inbound.GET("/:id", InboundHandler.Get)
			Inbound.POST("/", InboundHandler.Post)
			Inbound.POST("/:id", InboundHandler.Post)
			Inbound.PUT("/", InboundHandler.Put)
			Inbound.PUT("/:id", InboundHandler.Put)
			Inbound.PATCH("/", InboundHandler.Put)
			Inbound.PATCH("/:id", InboundHandler.Put)
			Inbound.DELETE("/", InboundHandler.Delete)
			Inbound.DELETE("/:id", InboundHandler.Delete)
		}

		dahuaNVR := api.Group("/dahua-nvr")
		{
			dahuaNVR.POST("/face-recognition", dahuaNVRHandler.ReceiveFaceRecognitionEvent)
			dahuaNVR.POST("/license-plate-recognition", dahuaNVRHandler.ReceiveLicensePlateRecognitionEvent)
		}

		dahuaCameraFaceRecognition := api.Group("/dahua/camera/face-recognition")
		{

			dahuaCameraFaceRecognition.POST("/data/", dahuaCameraFaceRecognitionHandler.ReceiveFaceRecognitionEvent)
			dahuaCameraFaceRecognition.POST("/picture/", dahuaCameraFaceRecognitionHandler.ReceiveFaceRecognitionImage)
		}

		webhook := api.Group("/webhook")
		{
			webhook.POST("/by-pass", webhookHandler.ByPassPost)
		}

	}

	return router
}
