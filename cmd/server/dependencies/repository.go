package dependencies

import (
	"github.com/phamtuanha21092003/mep-api-core/app/base"
	"github.com/phamtuanha21092003/mep-api-core/app/repository"
	"github.com/phamtuanha21092003/mep-api-core/platform/database"
	"github.com/phamtuanha21092003/mep-api-core/platform/logger"
)

type Repositories struct {
	userRepo repository.IUserRepository

	roleRepo repository.IRoleRepository
}

func InitRepositories(db *database.SqlxDatabase, logger *logger.Logger, transactionManagerSqlx base.ITransactionManagerSqlx) *Repositories {
	return &Repositories{
		userRepo: repository.NewUserRepository(db, logger, transactionManagerSqlx),

		roleRepo: repository.NewRoleRepository(db, logger, transactionManagerSqlx),
	}
}
