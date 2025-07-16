package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type (
	User struct {
		ID          uuid.UUID      `db:"id"           json:"id"`
		Password    string         `db:"password"     json:"-"`
		CreatedAt   time.Time      `db:"created_at"   json:"created_at"`
		UpdatedAt   time.Time      `db:"updated_at"   json:"updated_at"`
		DeletedAt   *time.Time     `db:"deleted_at"   json:"deleted_at,omitempty"`
		LastLogin   *time.Time     `db:"last_login"   json:"last_login,omitempty"`
		IsSuperuser bool           `db:"is_superuser" json:"is_superuser"`
		Username    string         `db:"username"     json:"username"`
		FirstName   string         `db:"first_name"   json:"first_name"`
		LastName    string         `db:"last_name"    json:"last_name"`
		Email       string         `db:"email"        json:"email"`
		IsActive    bool           `db:"is_active"    json:"is_active"`
		RoleID      sql.NullString `db:"role_id"      json:"role_id"`
	}
)
