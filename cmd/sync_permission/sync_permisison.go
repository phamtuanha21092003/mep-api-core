package sync_permission

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/phamtuanha21092003/mep-api-core/platform/database"
)

func SyncPermission(db *database.SqlxDatabase, routes gin.RoutesInfo) {
	for _, route := range routes {
		fmt.Println(route.Method, strings.Split(route.Path, "/"))
	}
}
