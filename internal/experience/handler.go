package experience

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListExperiencesHandler(c *gin.Context) {

	experiences, err := ListExperiences()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, experiences)

}
