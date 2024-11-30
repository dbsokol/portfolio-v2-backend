package choicefield

import (
	"fmt"
	"portfolio/internal/api/choices"
	"portfolio/internal/api/exception"
	"strings"

	"github.com/gin-gonic/gin"
)

func TextChoiceField(c *gin.Context, value string, fieldName string, choiceSet choices.TextChoices) *string {

	if value == "" {
		return nil
	}

	choice, err := choiceSet.Validate(value)
	if err != nil {
		apiException := exception.APIException{
			Name:   fieldName,
			Detail: fmt.Sprintf("field must be one of: %s", strings.Join(choiceSet.GetPublicList(), ", ")),
		}
		exception.AddAPIException(c, apiException)
		return nil
	}

	return &choice.Private
}

func IntChoiceField(c *gin.Context, value string, fieldName string, choiceSet choices.IntChoices) *int {

	if value == "" {
		return nil
	}

	choice, err := choiceSet.Validate(value)
	if err != nil {
		apiException := exception.APIException{
			Name:   fieldName,
			Detail: fmt.Sprintf("field must be one of: %s", strings.Join(choiceSet.GetPublicList(), ", ")),
		}
		exception.AddAPIException(c, apiException)
		return nil
	}

	return &choice.Private
}
