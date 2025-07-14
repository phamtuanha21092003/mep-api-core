package base

import (
	"context"
	"fmt"
	"strings"

	"github.com/phamtuanha21092003/mep-api-core/platform/database"
)

// Could not use base repo to query
type (
	IBaseRepositorySqlx[T any, ID any] interface {
		Paging(ctx context.Context, condition any, pagingInput PagingDto) ([]T, int64, error)

		GetById(ctx context.Context, id ID) (*T, error)
	}

	BaseRepositorySqlx[T any, ID any] struct {
		tableName string

		Sqlx *database.SqlxDatabase

		TransactionManager ITransactionManagerSqlx
	}
)

func NewBaseRepositorySqlx[T, ID any](db *database.SqlxDatabase, tableName string, transactionManagerSqlx ITransactionManagerSqlx) *BaseRepositorySqlx[T, ID] {
	repo := &BaseRepositorySqlx[T, ID]{Sqlx: db, tableName: tableName, TransactionManager: transactionManagerSqlx}
	// check compile-time with type assertion
	var _ IBaseRepositorySqlx[T, ID] = repo

	return repo
}

func (repo *BaseRepositorySqlx[T, ID]) GetTableName() string {
	return repo.tableName
}

// TODO: validate condition pre query
// TODO: thinking solution of join
func (repo *BaseRepositorySqlx[T, ID]) Paging(
	ctx context.Context,
	condition any,
	pagingInput PagingDto,
) ([]T, int64, error) {
	var (
		items []T
		total int64
	)

	if repo.tableName == "" {
		return nil, 0, fmt.Errorf("table name is not set")
	}

	if pagingInput.PageSize <= 0 {
		pagingInput.PageSize = 10
	}
	if pagingInput.PageNum <= 0 {
		pagingInput.PageNum = 1
	}

	offset := (pagingInput.PageNum - 1) * pagingInput.PageSize
	table := repo.tableName

	whereClause := ""
	params := map[string]interface{}{}

	if condition != nil {
		// Marshal condition to map[string]interface{} via JSON
		condMap, ok := convertToMap(condition)
		if ok && len(condMap) > 0 {
			var filters []string
			for key, val := range condMap {
				paramKey := key
				filters = append(filters, fmt.Sprintf("%s = :%s", key, paramKey))
				params[paramKey] = val
			}
			whereClause = "WHERE " + strings.Join(filters, " AND ")
		}
	}

	// --- Query paginated items ---
	query := fmt.Sprintf(
		"SELECT * FROM %s %s LIMIT :limit OFFSET :offset",
		table,
		whereClause,
	)
	params["limit"] = pagingInput.PageSize
	params["offset"] = offset

	stmt, err := repo.Sqlx.DB.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, 0, fmt.Errorf("prepare query failed: %w", err)
	}
	defer stmt.Close()

	if err := stmt.SelectContext(ctx, &items, params); err != nil {
		return nil, 0, fmt.Errorf("select paginated data failed: %w", err)
	}

	countQuery := fmt.Sprintf("SELECT COUNT(1) FROM %s %s", table, whereClause)
	countStmt, err := repo.Sqlx.DB.PrepareNamedContext(ctx, countQuery)
	if err != nil {
		return nil, 0, fmt.Errorf("prepare count query failed: %w", err)
	}
	defer countStmt.Close()

	if err := countStmt.GetContext(ctx, &total, params); err != nil {
		return nil, 0, fmt.Errorf("count query failed: %w", err)
	}

	return items, total, nil
}

func (repo *BaseRepositorySqlx[T, ID]) GetById(
	ctx context.Context,
	id ID,
) (*T, error) {

	if repo.tableName == "" {
		return nil, fmt.Errorf("table name is not set")
	}

	table := repo.tableName

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = :id LIMIT 1", table)

	var entity T

	if err := repo.Sqlx.DB.GetContext(ctx, &entity, query, id); err != nil {
		return nil, fmt.Errorf("Query failed: %w", err)
	}

	return &entity, nil
}
