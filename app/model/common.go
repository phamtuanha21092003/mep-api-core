package model

import "github.com/golang-jwt/jwt/v5"

type (
	UserClaims struct {
		jwt.RegisteredClaims
		Sub          string `json:"sub"`
		Email        string `json:"email"`
		Username     string `json:"username"`
		RoleID       string `json:"role_id"`
		TokenVersion int    `json:"token_version"`
	}
)
