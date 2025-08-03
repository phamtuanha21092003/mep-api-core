package common

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func ValidateExtractToken(c *gin.Context) (string, error) {
	bearToken := c.Request.Header.Get("Authorization")
	if bearToken == "" {
		return "", errors.New("Access token is missing or invalid")
	}

	var tokenAuth string
	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		tokenAuth = onlyToken[1]
	}
	if len(strings.TrimSpace(tokenAuth)) == 0 {
		return tokenAuth, errors.New("Access token is missing or invalid")
	}

	return tokenAuth, nil
}
