package dependencies

import (
	"github.com/phamtuanha21092003/mep-api-core/app/service"
	"github.com/phamtuanha21092003/mep-api-core/platform/logger"
)

type Services struct {
	userSer service.IUserService

	tokenSer service.ITokenService

	roleSer service.IRoleService

	tusSer service.ITusService
}

func InitServices(repo *Repositories, logger *logger.Logger) *Services {
	tokenSvc := service.NewTokenService(logger)

	return &Services{
		userSer: service.NewUserService(repo.userRepo, tokenSvc, logger),

		tokenSer: tokenSvc,

		roleSer: service.NewRoleService(repo.roleRepo, logger),

		tusSer: service.NewTusService(),
	}
}
