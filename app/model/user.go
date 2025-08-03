package model

import (
	"time"

	"github.com/google/uuid"

	"github.com/phamtuanha21092003/mep-api-core/app/base"
)

type (
	User struct {
		base.BaseModel[uuid.UUID]
		Password     string     `db:"password"     	json:"-"`
		DeletedAt    *time.Time `db:"deleted_at"   	json:"deleted_at,omitempty"`
		LastLogin    *time.Time `db:"last_login"   	json:"last_login,omitempty"`
		IsSuperuser  bool       `db:"is_superuser" 	json:"is_superuser"`
		Username     string     `db:"username"     	json:"username"`
		FirstName    string     `db:"first_name"   	json:"first_name"`
		LastName     string     `db:"last_name"    	json:"last_name"`
		Email        string     `db:"email"        	json:"email"`
		IsActive     bool       `db:"is_active"    	json:"is_active"`
		RoleID       *uuid.UUID `db:"role_id"      	json:"role_id"`
		TokenVersion int        `db:"token_version" json:"token_version"`
	}
)
