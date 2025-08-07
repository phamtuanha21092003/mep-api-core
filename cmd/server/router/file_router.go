package router

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/phamtuanha21092003/mep-api-core/cmd/server/dependencies"
)

func FileRouter(app *gin.Engine, contr dependencies.Controllers) {
	fileGroup := app.Group("/api/v1/files", contr.UserContr.VerifyPermissions([]string{"file:update"}))

	fileGroup.Any("/tus/*path", func(c *gin.Context) {
		fmt.Println(c.Request.URL.RawPath, " path ", c.Request.URL.Path)

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS")

		c.Writer.Header().Set("Access-Control-Expose-Headers", "Upload-Offset, Location, Upload-Length, Tus-Version, Tus-Resumable, Tus-Max-Size, Tus-Extension, Upload-Metadata, Upload-Defer-Length, Upload-Concat")

		c.Next()
	},
		contr.TusContr.Handler(),
	)

}
