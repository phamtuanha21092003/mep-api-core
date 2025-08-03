package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/phamtuanha21092003/mep-api-core/app/repository"
	"github.com/phamtuanha21092003/mep-api-core/pkg/utils"
	"github.com/phamtuanha21092003/mep-api-core/platform/logger"
)

type (
	IRoleService interface {
		VerifyPermission(ctx context.Context, roleID uuid.UUID, permissions []string) (bool, error)
	}

	RoleService struct {
		logger *logger.Logger

		roleRepo repository.IRoleRepository
	}
)

func NewRoleService(roleRepo repository.IRoleRepository, logger *logger.Logger) IRoleService {
	return &RoleService{roleRepo: roleRepo, logger: logger}
}

func (roleSer *RoleService) VerifyPermission(ctx context.Context, roleID uuid.UUID, permissions []string) (bool, error) {
	return roleSer.roleRepo.IsHavePermission(ctx, roleID, utils.BuildPermissionHierarchy(permissions))
}
