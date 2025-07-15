package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/phamtuanha21092003/mep-api-core/app/base"
	"github.com/phamtuanha21092003/mep-api-core/app/dto"
	"github.com/phamtuanha21092003/mep-api-core/app/service"
	"github.com/phamtuanha21092003/mep-api-core/pkg/utils"
)

type IUserController interface {
	Register() gin.HandlerFunc
}

type UserController struct {
	userSer service.IUserService
}

func NewUserController(userSer service.IUserService) IUserController {
	return &UserController{userSer: userSer}
}

func (userContr *UserController) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		dto, ok := utils.GetBodyRequest[dto.RegisterUserDto](c)
		if !ok {
			return
		}

		userSer := userContr.userSer

		userId, err := userSer.Register(c.Request.Context(), dto)
		if err != nil {
			c.Error(base.BadRequest(err))
			return
		}

		c.JSON(200, gin.H{
			"userId": userId,
		})
	}
}
