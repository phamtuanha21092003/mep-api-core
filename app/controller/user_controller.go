package controller

import (
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/phamtuanha21092003/mep-api-core/app/base"
	"github.com/phamtuanha21092003/mep-api-core/app/dto"
	"github.com/phamtuanha21092003/mep-api-core/app/service"
	"github.com/phamtuanha21092003/mep-api-core/pkg/utils"
)

type IUserController interface {
	Register() gin.HandlerFunc

	Login() gin.HandlerFunc

	Refresh() gin.HandlerFunc
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

func (userContr *UserController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		dto, ok := utils.GetBodyRequest[dto.LoginUserDto](c)
		if !ok {
			return
		}

		userSer := userContr.userSer

		accessToken, refreshToken, err := userSer.Login(c.Request.Context(), dto)
		if err != nil {
			c.Error(base.BadRequest(errors.New("Email or password is invalid!")))
			return
		}

		c.SetCookie(
			"refresh_token",
			refreshToken,
			60*60*24*30,
			"/",
			"",
			true,
			true,
		)

		c.JSON(200, gin.H{
			"access_token": accessToken,
		})
	}
}

func (userContr *UserController) Refresh() gin.HandlerFunc {
	return func(c *gin.Context) {
		refreshToken, err := c.Cookie("refresh_token")
		if err != nil {
			c.Error(base.BadRequest(errors.New("refresh token not found")))
			return
		}

		userSer := userContr.userSer
		accessToken, refreshToken, err := userSer.Refresh(c.Request.Context(), refreshToken)
		if err != nil {
			c.Error(base.BadRequest(err))
			return
		}

		c.SetCookie(
			"refresh_token",
			refreshToken,
			60*60*24*30,
			"/",
			"",
			true,
			true,
		)

		c.JSON(200, gin.H{
			"access_token": accessToken,
		})
	}
}
