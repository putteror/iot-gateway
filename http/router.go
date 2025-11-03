package http

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/putteror/iot-gateway/http/handler"
)

func NewRouter(
	defaultHandler *handler.DefaultHandler,
	InboundHandler *handler.InboundHandler,
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

	}

	return router
}
