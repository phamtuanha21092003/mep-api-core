package model

import (
	"github.com/google/uuid"

	"github.com/phamtuanha21092003/mep-api-core/app/base"
)

type RoleModel struct {
	base.BaseModel[uuid.UUID]
}
