package router

import (
	"github.com/gin-gonic/gin"

	"github.com/phamtuanha21092003/mep-api-core/cmd/server/dependencies"
)

func FileRouter(app *gin.Engine, contr dependencies.Controllers) {
	fileGroup := app.Group("/api/v1/files", contr.UserContr.VerifyPermissions([]string{"file:update"}))

	handler := contr.TusContr.Handler()

	fileGroup.POST("/tus/", gin.WrapF(handler.PostFile))
	fileGroup.HEAD("/tus/:id", gin.WrapF(handler.HeadFile))
	fileGroup.PATCH("/tus/:id", gin.WrapF(handler.PatchFile))
	fileGroup.GET("/tus/:id", gin.WrapF(handler.GetFile))
}
