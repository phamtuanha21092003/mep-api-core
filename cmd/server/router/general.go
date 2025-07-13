package router

import (
	"github.com/gin-gonic/gin"

	"github.com/phamtuanha21092003/mep-api-core/app/base"
	"github.com/phamtuanha21092003/mep-api-core/platform/database"
)

func GeneralRouter(app *gin.Engine) {
	app.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong from mep-core-api",
			"status":  "/h34l7h",
		})
	})

	app.GET("/h34l7h", func(c *gin.Context) {
		if err := database.SqlxConn.DB.Ping(); err != nil {
			c.Error(base.InternalServerError(err))
			return
		}
		c.JSON(200, gin.H{
			"message":   "Health Check",
			"db_online": true,
		})
	})
}
