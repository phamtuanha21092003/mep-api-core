package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/phamtuanha21092003/mep-api-core/app/base"
	"github.com/phamtuanha21092003/mep-api-core/pkg/config"
)

const (
	INTERNAL_SERVER_ERROR = "internal_server_error"
)

// GinMiddleware is apply to engine gin
func GinMiddleware(app *gin.Engine, config *config.Config) {
	// apply middleware cors
	app.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH, HEAD")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Upload-Offset, Location, Upload-Length, Tus-Version, Tus-Resumable, Tus-Max-Size, Tus-Extension, Upload-Metadata, Upload-Defer-Length, Upload-Concat")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	if config.GetBool("MIDDLEWARE_GIN_LOGGING_ENABLED") {
		app.Use(
			gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
				return fmt.Sprintf("[gateway] %s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
					param.ClientIP,
					param.TimeStamp.Format(time.RFC1123),
					param.Method,
					param.Path,
					param.Request.Proto,
					param.StatusCode,
					param.Latency,
					param.Request.UserAgent(),
					param.ErrorMessage,
				)
			}),
		)
	}

	if config.GetBool("MIDDLEWARE_GIN_RECOVER_ENABLED") {
		app.Use(gin.Recovery())
	}

	// TODO: add logging to this middleware
	// apply middleware error handling
	app.Use(func(c *gin.Context) {
		c.Next()

		last := c.Errors.Last()
		if last == nil || last.Err == nil {
			return
		}

		var appErr *base.AppError
		err := last.Err

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
