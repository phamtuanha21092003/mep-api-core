package router

import (
	"github.com/gin-gonic/gin"
)

func EventRouter(app *gin.Engine) {
	eventGroup := app.Group("/api/v1/event")

	eventGroup.GET("/:eventId")
	eventGroup.POST("/")
}
