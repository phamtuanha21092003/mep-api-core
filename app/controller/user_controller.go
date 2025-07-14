package controller

import "github.com/phamtuanha21092003/mep-api-core/app/service"

type IUserController interface {
}

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) IUserController {
	return &UserController{userService: userService}
}
