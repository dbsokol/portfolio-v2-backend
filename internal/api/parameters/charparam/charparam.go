package charparam

import (
	"fmt"
	"portfolio/internal/api/exception"

	"github.com/gin-gonic/gin"
)

func CharParam(c *gin.Context, name string, maxLength int) *string {
	value := c.Query(name)

	if len(value) > maxLength {
		apiException := exception.APIException{
			Name:   name,
			Detail: fmt.Sprintf("field must be less than %d characters", maxLength),
		}
		exception.AddAPIException(c, apiException)
		return nil
	}

	return &value
}
