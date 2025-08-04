package router

import (
	"github.com/gin-gonic/gin"

	"github.com/phamtuanha21092003/mep-api-core/cmd/server/dependencies"
)

func FileRouter(app *gin.Engine, contr dependencies.Controllers) {
	fileGroup := app.Group("/api/v1/files", contr.UserContr.VerifyPermissions([]string{"file:update"}))

	fileGroup.Any("/tus/*path", contr.TusContr.Handler())
}
