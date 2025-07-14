package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/phamtuanha21092003/mep-api-core/app/dto"
	"github.com/phamtuanha21092003/mep-api-core/pkg/utils"
)

func UserRouter(app *gin.Engine) {
	userGroup := app.Group("/api/v1/users")

	// TODO: move this to controller
	userGroup.POST("/register", func(c *gin.Context) {

		dto, ok := utils.GetBodyRequest[dto.RegisterUserDto](c)
		if !ok {
			return
		}

		fmt.Println(dto)

		c.JSON(200, gin.H{
			"test": "test",
		})
	})
}
