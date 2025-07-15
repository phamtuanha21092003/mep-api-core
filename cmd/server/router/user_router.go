package router

import (
	"github.com/gin-gonic/gin"

	"github.com/phamtuanha21092003/mep-api-core/app/controller"
)

func UserRouter(app *gin.Engine, userController controller.IUserController) {
	userGroup := app.Group("/api/v1/users")

	userGroup.POST("/register", userController.Register())
}
