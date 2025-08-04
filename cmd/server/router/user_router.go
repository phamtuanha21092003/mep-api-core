package router

import (
	"github.com/gin-gonic/gin"
)

func UserRouter(app *gin.Engine) {
	userGroup := app.Group("/api/v1/users")

	userGroup.GET("/me")
}
