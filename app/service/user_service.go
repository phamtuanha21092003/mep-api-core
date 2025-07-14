package service

import (
	"github.com/phamtuanha21092003/mep-api-core/app/repository"
	"github.com/phamtuanha21092003/mep-api-core/platform/logger"
)

type IUserService interface{}

type UserService struct {
	userRepo repository.IUserRepository
	logger   *logger.Logger
}

func NewUserService(userRepo repository.IUserRepository, logger *logger.Logger) IUserService {
	return &UserService{userRepo: userRepo, logger: logger}
}
