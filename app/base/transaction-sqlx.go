package base

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// ITransactionManagerSqlx wraps an atomic unit of work.
type ITransactionManagerSqlx interface {
	Do(ctx context.Context, fn func(*sqlx.Tx) (interface{}, error)) (interface{}, error)
}

// txManagerSqlx is the concrete sqlx implementation.
type txManagerSqlx struct {
	db *sqlx.DB
}

// NewTxManager creates a Transaction
func NewTxManager(db *sqlx.DB) ITransactionManagerSqlx {
	return &txManagerSqlx{db: db}
}

// Do implements Transaction using BEGIN … COMMIT / ROLLBACK.
func (m *txManagerSqlx) Do(ctx context.Context, fn func(*sqlx.Tx) (interface{}, error)) (interface{}, error) {
	tx, err := m.db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	// Defer handles both panic and error cases.
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	data, err := fn(tx)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return data, nil
}
