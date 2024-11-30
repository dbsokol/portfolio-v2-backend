package dateparam

import (
	"portfolio/internal/api/exception"
	"time"

	"github.com/gin-gonic/gin"
)

func DateParam(c *gin.Context, name string) *time.Time {
	// Get the query parameter value
	value := c.Query(name)

	// If the value is empty, return nil (no error)
	if value == "" {
		return nil
	}

	// Try to parse the value as a date
	dateValue, err := time.Parse("2006-01-02", value)
	if err != nil {
		apiException := exception.APIException{
			Name:   name,
			Detail: "field must be a valid date (YYYY-MM-DD)",
		}
		exception.AddAPIException(c, apiException)
		return nil

	}

	// Return the parsed date value
	return &dateValue
}
