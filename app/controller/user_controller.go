package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/phamtuanha21092003/mep-api-core/app/base"
	"github.com/phamtuanha21092003/mep-api-core/app/common"
	"github.com/phamtuanha21092003/mep-api-core/app/dto"
	"github.com/phamtuanha21092003/mep-api-core/app/service"
	"github.com/phamtuanha21092003/mep-api-core/pkg/utils"
)

type IUserController interface {
	Register() gin.HandlerFunc

	Login() gin.HandlerFunc

	Refresh() gin.HandlerFunc

	VerifyPermissions(permissions []string) gin.HandlerFunc
}

type UserController struct {
	userSer service.IUserService

	tokenSer service.ITokenService

	roleSer service.IRoleService
}

func NewUserController(userSer service.IUserService, tokenSer service.ITokenService, roleSer service.IRoleService) IUserController {
	return &UserController{userSer: userSer, tokenSer: tokenSer, roleSer: roleSer}
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

func (userContr *UserController) VerifyPermissions(permissions []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := common.ValidateExtractToken(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "forbidden",
				"error":   err.Error(),
			})
			return
		}

		claim, err := userContr.tokenSer.VerifyUserToken(accessToken, utils.JWT_ACCESS_TOKEN)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "forbidden",
				"error":   err.Error(),
			})
			return
		}

		if len(permissions) == 0 || claim.IsSuperuser {
			c.Next()
		}

		if claim.RoleID == nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "forbidden",
				"error":   "Permission is not allow",
			})
			return
		}

		isHavePermission, err := userContr.roleSer.VerifyPermission(c.Request.Context(), *claim.RoleID, permissions)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "forbidden",
				"error":   err.Error(),
			})
			return
		}

		if isHavePermission == false {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "forbidden",
				"error":   "Permission is not allow",
			})
			return
		}

		c.Next()
	}
}
