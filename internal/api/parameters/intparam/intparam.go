package intparam

import (
	"portfolio/internal/api/exception"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IntParam(c *gin.Context, name string) *int {
	// Get the query parameter value
	value := c.Query(name)

	// If the value is empty, return nil (no error)
	if value == "" {
		return nil
	}

	// Try to parse the value as an integer
	intValue, err := strconv.Atoi(value)
	if err != nil {
		apiException := exception.APIException{
			Name:   name,
			Detail: "field must be an integer",
		}
		exception.AddAPIException(c, apiException)
		return nil
	}

	// Return the parsed integer value
	return &intValue
}
