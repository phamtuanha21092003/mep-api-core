package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/phamtuanha21092003/mep-api-core/app/base"
)

const (
	INTERNAL_SERVER_ERROR = "internal_server_error"
)

// GinMiddleware is apply to engine gin
func GinMiddleware(app *gin.Engine) {
	// apply middleware cors
	app.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// TODO: add logging to this middleware
	// apply middleware error handling
	app.Use(func(c *gin.Context) {
		c.Next()

		var appErr *base.AppError
		err := c.Errors.Last().Err

		if errors.As(err, &appErr) {
			appErr := err.(*base.AppError)
			c.AbortWithStatusJSON(appErr.Code, gin.H{
				"message": appErr.Message,
				"error":   appErr.Err.Error(),
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": INTERNAL_SERVER_ERROR,
			"error":   err.Error(),
		})
	})
}
