package model

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type (
	UserClaims struct {
		jwt.RegisteredClaims
		Sub          string     `json:"sub"`
		Email        string     `json:"email"`
		Username     string     `json:"username"`
		RoleID       *uuid.UUID `json:"role_id,omitempty"`
		TokenVersion int        `json:"token_version"`
		IsSuperuser  bool       `json:"is_superuser"`
	}
)
