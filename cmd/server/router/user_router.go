package router

import (
	"github.com/gin-gonic/gin"
)

func UserRouter(app *gin.Engine) {
	userGroup := app.Group("api/v1/users")

	userGroup.POST("/register", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"test": "test",
		})
	})
}
