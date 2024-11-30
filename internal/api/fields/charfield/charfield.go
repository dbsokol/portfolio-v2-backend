package charfield

import (
	"portfolio/internal/api/exception"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CharField(c *gin.Context, char string, fieldName string, maxLength int) *string {
	if char == "" {
		return nil
	}

	if len(char) > maxLength {
		apiException := exception.APIException{
			Name:   fieldName,
			Detail: "field must be less than " + strconv.Itoa(maxLength) + " characters",
		}
		exception.AddAPIException(c, apiException)
		return nil
	}
	return &char
}
