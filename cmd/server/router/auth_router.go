package router

import (
	"github.com/gin-gonic/gin"

	"github.com/phamtuanha21092003/mep-api-core/app/controller"
)

func AuthRouter(app *gin.Engine, userController controller.IUserController) {
	authGroup := app.Group("/api/v1/auth")

	authGroup.POST("/register", userController.Register())
	authGroup.POST("/login", userController.Login())
}
