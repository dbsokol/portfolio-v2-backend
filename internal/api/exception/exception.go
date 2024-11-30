package exception

import (
	"github.com/gin-gonic/gin"
)

type APIException struct {
	Name   string `json:"name"`
	Detail string `json:"detail"`
}

func AddAPIException(c *gin.Context, apiException APIException) {
	if errs, exists := c.Get("APIExceptions"); exists {
		errorList := errs.(*[]APIException)
		*errorList = append(*errorList, apiException)
	} else {
		c.Set("APIExceptions", &[]APIException{apiException})
	}
}

func GetAPIExceptions(c *gin.Context) []APIException {
	if errs, exists := c.Get("APIExceptions"); exists {
		return *errs.(*[]APIException)
	}
	return []APIException{}
}
