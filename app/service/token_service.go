package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/phamtuanha21092003/mep-api-core/app/model"
	"github.com/phamtuanha21092003/mep-api-core/pkg/config"
	"github.com/phamtuanha21092003/mep-api-core/pkg/utils"
	"github.com/phamtuanha21092003/mep-api-core/platform/logger"
)

type (
	ITokenService interface {
		CreateUserToken(claim *model.UserClaims, tokenType utils.TokenType) (string, error)
	}

	TokenService struct {
		logger *logger.Logger
	}
)

func NewTokenService(logger *logger.Logger) ITokenService {
	return &TokenService{logger: logger}
}

func (tokenSer *TokenService) CreateUserToken(claim *model.UserClaims, tokenType utils.TokenType) (string, error) {
	exp := time.Now().Add(time.Minute * time.Duration(config.AppCfg().JWTSecretRefreshExpire))
	secret := []byte(config.AppCfg().JWTRefreshTokenSecretKey)

	if tokenType == utils.JWT_ACCESS_TOKEN {
		exp = time.Now().Add(time.Minute * time.Duration(config.AppCfg().JWTSecretExpire))
		secret = []byte(config.AppCfg().JWTSecretKey)
	}
	claim.ExpiresAt = jwt.NewNumericDate(exp)
	claim.IssuedAt = jwt.NewNumericDate(time.Now())

	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := tok.SignedString(secret)
	if err != nil {
		return "", err
	}

	return token, nil
}
