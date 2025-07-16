package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/phamtuanha21092003/mep-api-core/app/dto"
	"github.com/phamtuanha21092003/mep-api-core/app/repository"
	"github.com/phamtuanha21092003/mep-api-core/pkg/utils"
	"github.com/phamtuanha21092003/mep-api-core/platform/logger"
)

type IUserService interface {
	Register(ctx context.Context, payload *dto.RegisterUserDto) (any, error)
	Login(ctx context.Context, payload *dto.LoginUserDto) (string, string, error)
}

type UserService struct {
	userRepo repository.IUserRepository
	logger   *logger.Logger
}

func NewUserService(userRepo repository.IUserRepository, logger *logger.Logger) IUserService {
	return &UserService{userRepo: userRepo, logger: logger}
}

func (userSer *UserService) Register(ctx context.Context, payload *dto.RegisterUserDto) (any, error) {
	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		return "", err
	}

	payload.Password = hashedPassword
	payload.ID = uuid.NewString()

	return userSer.userRepo.Register(ctx, payload)
}

// TODO: generate access and refresh token use jwt
func (userSer *UserService) Login(ctx context.Context, payload *dto.LoginUserDto) (string, string, error) {
	user, err := userSer.userRepo.GetUserLogin(ctx, payload.Email)
	if err != nil {
		return "", "", err
	}

	err = utils.CheckPasswordHash(payload.Password, user.Password)
	if err != nil {
		return "", "", err
	}

	return user.ID.String(), user.Password, nil
}
