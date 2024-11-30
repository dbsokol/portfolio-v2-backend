package contact

import (
	"net/http"
	"portfolio/internal/api/exception"
	"portfolio/internal/api/fields"
	"portfolio/internal/api/pagination"
	"portfolio/internal/api/parameters"

	"github.com/gin-gonic/gin"
)

type CreateContactRequest struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	URL   string `json:"url"` // Optional
}

// CreateContactHandler creates a new contact
// @Summary Create a new contact
// @Description Add a new contact to the database
// @Tags contacts
// @Param contact body CreateContactRequest true "Contact data"
// @Produce json
// @Success 201 {object} Contact
// @Router /contacts [post]
func CreateContactHandler(c *gin.Context) {
	var req CreateContactRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request", "detail": err.Error()})
		return
	}

	name := fields.CharField(c, req.Name, "name", 100)
	value := fields.CharField(c, req.Value, "value", 100)
	url := fields.CharField(c, req.URL, "url", 100)

	contact, err := CreateContact(*name, *value, url)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contact)
}

// ListContactsHandler lists all contacts
// @Summary List all contacts
// @Description List all contacts in the database
// @Tags contacts
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Param ordering query []string false "Ordering" Enums(CREATED_AT, -CREATED_AT, NAME, -NAME)
// @Param nameIContains query string false "Name insensitive contains"
// @Produce json
// @Success 200 {object} Contact
// @Router /contacts [get]
func ListContactsHandler(c *gin.Context) {
	limit := parameters.IntParam(c, "limit")
	offset := parameters.IntParam(c, "offset")
	ordering := parameters.TextChoiceListParam(c, "ordering", ContactOrderingChoices)
	nameIContains := parameters.CharParam(c, "nameIContains", 100)

	errors := exception.GetAPIExceptions(c)

	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	contacts, count := ListContacts(
		limit,
		offset,
		ordering,
		nameIContains,
	)

	response := pagination.GetPaginatedResponse(
		c,
		limit,
		offset,
		count,
		contacts,
	)

	// check for pagination specific errors
	errors = exception.GetAPIExceptions(c)
	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	c.JSON(http.StatusOK, response)
}
