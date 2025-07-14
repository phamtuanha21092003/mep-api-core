package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/phamtuanha21092003/mep-api-core/app/base"
	"github.com/phamtuanha21092003/mep-api-core/app/dto"
	"github.com/phamtuanha21092003/mep-api-core/app/model"
	"github.com/phamtuanha21092003/mep-api-core/platform/database"
	"github.com/phamtuanha21092003/mep-api-core/platform/logger"
)

type IUserRepository interface {
	base.IBaseRepositorySqlx[model.User, uuid.UUID]
	Register(ctx context.Context, payload *dto.RegisterUserDto) (string, error)
}

type UserRepository struct {
	*base.BaseRepositorySqlx[model.User, uuid.UUID]
	logger *logger.Logger
}

func NewUserRepository(db *database.SqlxDatabase, logger *logger.Logger, transactionManagerSqlx base.ITransactionManagerSqlx) IUserRepository {
	return &UserRepository{
		BaseRepositorySqlx: base.NewBaseRepositorySqlx[model.User, uuid.UUID](db, "user", transactionManagerSqlx),
		logger:             logger,
	}
}

func (userRepo *UserRepository) Register(ctx context.Context, payload *dto.RegisterUserDto) (string, error) {
	id, err := userRepo.TransactionManager.Do(ctx, func(tx *sqlx.Tx) (any, error) {
		query := `
			INSERT INTO "user" (id, email, username, password, first_name, last_name, is_superuser, is_active, created_at, updated_at)
			VALUES (:id, :email, :username, :password, :first_name, :last_name, :is_superuser, :is_active, :created_at, :updated_at)
			RETURNING id
		`

		err := tx.QueryRowxContext(ctx, query, payload).Scan(payload.ID)
		if err != nil {
			return "", err
		}

		return payload.ID, nil

	})

	return id.(string), err
}
