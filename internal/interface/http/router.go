package http

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/putteror/iot-gateway/internal/interface/http/handler"
	dahuahandler "github.com/putteror/iot-gateway/internal/interface/http/handler/thirdparty/dahua"
	hikvisionhandler "github.com/putteror/iot-gateway/internal/interface/http/handler/thirdparty/hikvision"
)

func NewRouter(
	thirdPartyDahuaHandler *dahuahandler.DahuaCameraFaceRecognitionHandler,
	thirdPartyHikvisionHandler *hikvisionhandler.HikvisionCameraEmergencyAlarmHandler,
	defaultHandler *handler.DebugHandler,
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

	thirdPartyApi := router.Group("/third-party")
	{
		dahuaApi := thirdPartyApi.Group("/dahua")
		{
			dahuaCameraApi := dahuaApi.Group("/camera")
			{
				dahuaCameraApi.POST("/face-recognition/:id", thirdPartyDahuaHandler.FaceRecognitionEvent)
				dahuaCameraApi.POST("/face-recognition/picture/:id", thirdPartyDahuaHandler.FaceRecognitionImageEvent)
			}
		}

		hikvisionApi := thirdPartyApi.Group("/hikvision")
		{
			hikvisionCameraApi := hikvisionApi.Group("/camera")
			{
				hikvisionCameraApi.POST("/emergency-alarm/:id/site/:siteId", thirdPartyHikvisionHandler.EmergencyAlarmEvent)
			}
		}
	}

	api := router.Group("/api")
	{

		// Access Control Device endpoints
		apiDebug := api.Group("/debug")
		{
			apiDebugPrint := apiDebug.Group("/print")
			{
				apiDebugPrint.GET("/", defaultHandler.ConsolePrint)
				apiDebugPrint.GET("/:id", defaultHandler.ConsolePrint)
				apiDebugPrint.POST("/", defaultHandler.ConsolePrint)
				apiDebugPrint.PUT("/:id", defaultHandler.ConsolePrint)
				apiDebugPrint.PATCH("/", defaultHandler.ConsolePrint)
				apiDebugPrint.DELETE("/:id", defaultHandler.ConsolePrint)
			}
			apiDebugPush := apiDebug.Group("/push")
			{
				apiDebugPush.GET("/", defaultHandler.PushRawData)
				apiDebugPush.POST("/", defaultHandler.PushRawData)
				apiDebugPush.PUT("/", defaultHandler.PushRawData)
				apiDebugPush.PATCH("/", defaultHandler.PushRawData)
				apiDebugPush.DELETE("/", defaultHandler.PushRawData)
			}

		}

	}

	return router
}
