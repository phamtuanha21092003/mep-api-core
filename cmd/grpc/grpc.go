package grpc

import (
	"github.com/phamtuanha21092003/mep-api-core/app/base"
	"github.com/phamtuanha21092003/mep-api-core/cmd/server/dependencies"
	"github.com/phamtuanha21092003/mep-api-core/platform/database"
	"github.com/phamtuanha21092003/mep-api-core/platform/logger"
)

type Grpc struct {
	sqlx         *database.SqlxDatabase
	logger       *logger.Logger
	repositories dependencies.Repositories
	services     dependencies.Services
}

func NewGrpc(db *database.SqlxDatabase) (*Grpc, error) {
	logger.SetUpLogger()
	logger := logger.GetLogger()

	transactionManager := base.NewTxManagerSqlx(db.DB)

	repos := dependencies.InitRepositories(db, logger, transactionManager)

	services := dependencies.InitServices(repos, logger)

	return &Grpc{
		sqlx:         db,
		logger:       logger,
		repositories: *repos,
		services:     *services,
	}, nil
}
