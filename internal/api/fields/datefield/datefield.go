package datefield

import (
	"portfolio/internal/api/exception"
	"time"

	"github.com/gin-gonic/gin"
)

func DateField(c *gin.Context, dateString string, fieldName string) *time.Time {
	if dateString == "" {
		return nil
	}

	date, err := time.Parse("2006-01-02", dateString)

	if err != nil {
		apiException := exception.APIException{
			Name:   fieldName,
			Detail: "field must be a valid date",
		}
		exception.AddAPIException(c, apiException)
		return nil
	}

	return &date
}
