package choicelistparam

import (
	"portfolio/internal/api/choices"
	"portfolio/internal/api/exception"
	"strings"

	"github.com/gin-gonic/gin"
)

func TextChoiceListParam(c *gin.Context, name string, choiceSet choices.TextChoices) []string {
	var validatedValues []string

	// Get the query parameter value
	value := c.Query(name)

	// If the value is empty, return an empty list (no error)
	if value == "" {
		return []string{}
	}

	valueList := strings.Split(value, ",")

	for _, v := range valueList {
		trimmed := strings.TrimSpace(v)

		// Validate against the Choices object
		choice, err := choiceSet.Validate(trimmed)
		if err != nil {
			apiException := exception.APIException{
				Name:   name,
				Detail: "field must be one of: " + strings.Join(choiceSet.GetPublicList(), ", "),
			}
			exception.AddAPIException(c, apiException)
			return nil
		}

		validatedValues = append(validatedValues, choice.Private)
	}

	return validatedValues
}

func IntChoiceListParam(c *gin.Context, name string, choiceSet choices.IntChoices) []int {
	var validatedValues []int

	// Get the query parameter value
	value := c.Query(name)

	// If the value is empty, return an empty list (no error)
	if value == "" {
		return []int{}
	}

	valueList := strings.Split(value, ",")

	for _, v := range valueList {
		trimmed := strings.TrimSpace(v)

		// Validate against the Choices object
		choice, err := choiceSet.Validate(trimmed)
		if err != nil {
			apiException := exception.APIException{
				Name:   name,
				Detail: "field must be one of: " + strings.Join(choiceSet.GetPublicList(), ", "),
			}
			exception.AddAPIException(c, apiException)
			return nil
		}

		validatedValues = append(validatedValues, choice.Private)
	}

	return validatedValues
}
