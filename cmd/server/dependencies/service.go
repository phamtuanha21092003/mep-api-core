package dependencies

import (
	"github.com/phamtuanha21092003/mep-api-core/app/service"
	"github.com/phamtuanha21092003/mep-api-core/platform/logger"
)

type Services struct {
	userService service.IUserService
}

func InitServices(repo *Repositories, logger *logger.Logger) *Services {
	return &Services{userService: service.NewUserService(repo.userRepo, logger)}
}
