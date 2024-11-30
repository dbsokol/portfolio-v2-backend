package skill

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"portfolio/internal/api/exception"
	"portfolio/internal/api/fields"
	"portfolio/internal/api/pagination"
	"portfolio/internal/api/parameters"
)

// ListSkillsHandler godoc
// @Summary List skills
// @Description Retrieve a list of skills from the database
// @Tags skills
// @Param limit query int false "Limit the number of results returned"
// @Param offset query int false "Offset the results returned"
// @Param startDateGTE query string false "Filter by start date greater than or equal to (YYYY-MM-DD)"
// @Param startDateLTE query string false "Filter by start date less than or equal to (YYYY-MM-DD)"
// @Param types query []string false "Filter by skill types (multi-select)" Enums(LANGUAGE, FRAMEWORK, CLOUD)
// @Param ordering query []string false "Order by fields (multi-select)" Enums(CREATED_AT, -CREATED_AT, NAME, -NAME, START_DATE, -START_DATE)
// @Produce json
// @Success 200 {array} Skill
// @Router /skills [get]
func ListSkillsHandler(c *gin.Context) {

	limit := parameters.IntParam(c, "limit")
	offset := parameters.IntParam(c, "offset")
	startDateGTE := parameters.DateParam(c, "startDateGTE")
	startDateLTE := parameters.DateParam(c, "startDateLTE")
	types := parameters.IntChoiceListParam(c, "types", SkillTypeChoices)
	ordering := parameters.TextChoiceListParam(c, "ordering", SkillOrderingChoices)

	// if there are errors, exit before executing any SQL queries
	errors := exception.GetAPIExceptions(c)
	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	skills, count := ListSkills(
		limit,
		offset,
		startDateGTE,
		startDateLTE,
		types,
		ordering,
	)

	response := pagination.GetPaginatedResponse(
		c,
		limit,
		offset,
		count,
		skills,
	)

	// check for pagination specific errors
	errors = exception.GetAPIExceptions(c)
	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	c.JSON(http.StatusOK, response)

}

type CreateSkillRequest struct {
	Name      string `json:"name" binding:"required"`
	StartDate string `json:"startDate" binding:"required"`
	Type      string `json:"type" binding:"required"`
}

// CreateSkillHandler godoc
// @Summary Create a new skill
// @Description Add a new skill to the database
// @Tags skills
// @Param skill body CreateSkillRequest true "Skill data"
// @Produce json
// @Success 201 {object} Skill
// @Router /skills [post]
func CreateSkillHandler(c *gin.Context) {
	var req CreateSkillRequest

	// Bind and validate request JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request", "detail": err.Error()})
		return
	}

	startDate := fields.DateField(c, req.StartDate, "startDate")
	name := fields.CharField(c, req.Name, "name", 100)
	skillType := fields.IntChoiceField(c, req.Type, "type", SkillTypeChoices)

	errors := exception.GetAPIExceptions(c)

	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	// Create Skill Object
	skill, err := CreateSkill(
		*name,
		*startDate,
		*skillType,
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error":  "failed to create skill",
				"detail": err.Error(),
			},
		)
		return
	}

	// Respond with Created Skill
	c.JSON(http.StatusCreated, skill)
}

// DeleteSkillHandler godoc
// @Summary Delete a skill
// @Description Delete a skill from the database
// @Tags skills
// @Param uuid path string true "Skill UUID"
// @Produce json
// @Success 204
// @Router /skills/{uuid}/ [delete]
func DeleteSkillHandler(c *gin.Context) {
	uuid := c.Param("uuid")

	err := DeleteSkill(uuid)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error":  "failed to delete skill",
				"detail": err.Error(),
			},
		)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

type PartialUpdateSkillRequest struct {
	Name      string `json:"name"`      // Optional
	StartDate string `json:"startDate"` // Optional
	Type      string `json:"type"`      // Optional
}

// PartialUpdateSkillHandler godoc
// @Summary Partially update a skill
// @Description Partially update a skill in the database
// @Tags skills
// @Param uuid path string true "Skill UUID"
// @Param skill body PartialUpdateSkillRequest true "Skill data"
// @Produce json
// @Success 200 {object} Skill
// @Router /skills/{uuid}/ [patch]
func PartialUpdateSkillHandler(c *gin.Context) {
	uuid := c.Param("uuid")

	skill, err := GetSkill(uuid)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "object with uuid=" + uuid + " not found",
			},
		)
		return
	}

	var req PartialUpdateSkillRequest

	// Bind and validate request JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request", "detail": err.Error()})
		return
	}

	startDate := fields.DateField(c, req.StartDate, "startDate")
	name := fields.CharField(c, req.Name, "name", 100)
	skillType := fields.IntChoiceField(c, req.Type, "type", SkillTypeChoices)

	errors := exception.GetAPIExceptions(c)

	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	// Create Skill Object
	skill, err = UpdateSkill(
		*skill,
		name,
		startDate,
		skillType,
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error":  "failed to update skill",
				"detail": err.Error(),
			},
		)
		return
	}

	// Respond with Updated Skill
	c.JSON(http.StatusOK, skill)
}
