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

	Register(ctx context.Context, payload *dto.RegisterUserDto) (any, error)

	GetUserLogin(ctx context.Context, email string) (*model.User, error)
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

func (userRepo *UserRepository) Register(ctx context.Context, payload *dto.RegisterUserDto) (any, error) {
	return userRepo.TransactionManager.Do(ctx, func(tx *sqlx.Tx) (any, error) {
		query := `
			INSERT INTO "user" (id, email, username, password, first_name, last_name, is_superuser, is_active)
			VALUES (:id, :email, :username, :password, :first_name, :last_name, false, false)
			RETURNING id;
		`

		stmt, err := tx.PrepareNamedContext(ctx, query)
		if err != nil {
			return "", err
		}
		defer stmt.Close()

		if err := stmt.GetContext(ctx, &payload.ID, payload); err != nil {
			return "", err
		}

		return payload.ID, nil
	})
}

func (userRepo *UserRepository) GetUserLogin(ctx context.Context, email string) (*model.User, error) {
	query := `SELECT u.id, u.email, u.username,  u.password, u.role_id, u.token_version FROM "user" u WHERE u.email = $1;`

	var user model.User
	if err := userRepo.Sqlx.DB.GetContext(ctx, &user, query, email); err != nil {
		return nil, err
	}

	return &user, nil
}
