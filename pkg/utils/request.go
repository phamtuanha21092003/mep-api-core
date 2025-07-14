package utils

import (
	"github.com/gin-gonic/gin"

	"github.com/phamtuanha21092003/mep-api-core/app/base"
)

func GetBodyRequest[T interface{}](c *gin.Context) (*T, bool) {
	var inputType T
	if err := c.ShouldBind(&inputType); err != nil {
		c.Error(base.BadRequest(err))
		return nil, false
	}

	return &inputType, true
}

func GetQueryRequest[T interface{}](c *gin.Context) (*T, bool) {
	var inputType T
	if err := c.ShouldBindQuery(&inputType); err != nil {
		c.Error(base.BadRequest(err))
		return nil, false
	}

	return &inputType, true
}
