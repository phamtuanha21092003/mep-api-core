package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"

	"github.com/phamtuanha21092003/mep-api-core/app/base"
	"github.com/phamtuanha21092003/mep-api-core/app/model"
	"github.com/phamtuanha21092003/mep-api-core/platform/database"
	"github.com/phamtuanha21092003/mep-api-core/platform/logger"
)

type IRoleRepository interface {
	base.IBaseRepositorySqlx[model.User, uuid.UUID]

	IsHavePermission(ctx context.Context, roleID uuid.UUID, permissions []string) (bool, error)
}

type RoleRepository struct {
	*base.BaseRepositorySqlx[model.User, uuid.UUID]
	logger *logger.Logger
}

func NewRoleRepository(db *database.SqlxDatabase, logger *logger.Logger, transactionManagerSqlx base.ITransactionManagerSqlx) IRoleRepository {
	return &RoleRepository{
		BaseRepositorySqlx: (*base.BaseRepositorySqlx[model.User, uuid.UUID])(base.NewBaseRepositorySqlx[model.RoleModel, uuid.UUID](db, "role", transactionManagerSqlx)),
		logger:             logger,
	}
}

func (roleRepo *RoleRepository) IsHavePermission(ctx context.Context, roleID uuid.UUID, permissions []string) (bool, error) {
	startAt := 2

	valuesJoinQuery, joinArgs, err := roleRepo.Sqlx.BuildSQLJoinValues(permissions, []string{}, &startAt)
	if err != nil {
		return false, err
	}

	query := fmt.Sprintf(`
		SELECT 1
		FROM role_permission rp
		JOIN permission p ON rp.permission_id = p.id
		JOIN (VALUES %s) AS pn(name) ON p.name = pn.name
		WHERE rp.role_id = $1
		LIMIT 1
	`, valuesJoinQuery)

	args := append([]any{roleID}, joinArgs...)
	var exists int
	err = roleRepo.Sqlx.DB.QueryRowContext(ctx, query, args...).Scan(&exists)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil

}
