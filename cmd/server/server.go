package server

import (
	"github.com/gin-gonic/gin"
	"github.com/phamtuanha21092003/mep-api-core/platform/database"
)

type Server struct {
	gin  *gin.Engine
	sqlx *database.SqlxDatabase
}
