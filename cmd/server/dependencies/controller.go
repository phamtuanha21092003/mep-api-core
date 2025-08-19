package dependencies

import "github.com/phamtuanha21092003/mep-api-core/app/controller"

type Controllers struct {
	UserContr controller.IUserController
}

func InitController(service *Services) *Controllers {
	return &Controllers{
		UserContr: controller.NewUserController(service.userSer, service.tokenSer, service.roleSer),
	}
}
