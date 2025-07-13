package repository

import (
	"github.com/google/uuid"

	"github.com/phamtuanha21092003/mep-api-core/app/base"
	"github.com/phamtuanha21092003/mep-api-core/app/model"
	"github.com/phamtuanha21092003/mep-api-core/platform/database"
	"github.com/phamtuanha21092003/mep-api-core/platform/logger"
)

type IUserRepository interface {
	base.IBaseRepositorySqlx[model.User, uuid.UUID]
}

type UserRepository struct {
	*base.BaseRepositorySqlx[model.User, uuid.UUID]
	logger *logger.Logger
}

func NewUserRepository(db *database.SqlxDatabase, logger *logger.Logger) IUserRepository {
	return &UserRepository{
		BaseRepositorySqlx: base.NewBaseRepositorySqlx[model.User, uuid.UUID](db, "user"),
		logger:             logger,
	}
}
