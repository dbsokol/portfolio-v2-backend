package education

import (
	"net/http"
	"portfolio/internal/api/exception"
	"portfolio/internal/api/fields"
	"portfolio/internal/api/pagination"
	"portfolio/internal/api/parameters"

	"github.com/gin-gonic/gin"
)

type CreateEducationRequest struct {
	Institution string `json:"institution" binding:"required"`
	Major       string `json:"major" binding:"required"`
	Degree      string `json:"degree" binding:"required"`
	StartDate   string `json:"startDate" binding:"required"`
	EndDate     string `json:"endDate" binding:"required"`
}

// CreateEducationHanlder creates a new education record
// @Summary Create a new education record
// @Description Add a new education record to the database
// @Tags educations
// @Param education body CreateEducationRequest true "Education data"
// @Produce json
// @Success 201 {object} Education
// @Router /educations [post]
func CreateEducationHandler(c *gin.Context) {
	var req CreateEducationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error":  "invalid request",
				"detail": err.Error(),
			},
		)
		return
	}

	institution := fields.CharField(c, req.Institution, "institution", 100)
	major := fields.CharField(c, req.Major, "major", 100)
	degree := fields.IntChoiceField(c, req.Degree, "degree", DegreeChoices)
	startDate := fields.DateField(c, req.StartDate, "startDate")
	endDate := fields.DateField(c, req.EndDate, "endDate")

	errors := exception.GetAPIExceptions(c)

	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	education, err := CreateEducation(
		*institution,
		*degree,
		*major,
		*startDate,
		*endDate,
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error":  "failed to create education",
				"detail": err.Error(),
			},
		)
		return
	}

	c.JSON(http.StatusCreated, education)

}

// DeleteEducationHandler deletes an education record
// @Summary Delete an education record
// @Description Soft delete an education record from the database
// @Tags educations
// @Param uuid path string true "Education UUID"
// @Produce json
// @Success 204
// @Router /educations/{uuid}/ [delete]
func DeleteEducationHandler(c *gin.Context) {
	uuid := c.Param("uuid")

	err := DeleteEducation(uuid)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error":  "failed to delete education",
				"detail": err.Error(),
			},
		)
		return
	}

	c.Status(http.StatusNoContent)
}

type PartialUpdateEducationRequest struct {
	Institution string `json:"institution"` // Optional
	Major       string `json:"major"`       // Optional
	Degree      string `json:"degree"`      // Optional
	StartDate   string `json:"startDate"`   // Optional
	EndDate     string `json:"endDate"`     // Optional
}

// PartialUpdateEducationHandler updates an education record
// @Summary Partially update an education record
// @Description Partially update an education record in the database
// @Tags educations
// @Param uuid path string true "Education UUID"
// @Param education body PartialUpdateEducationRequest true "Education data"
// @Produce json
// @Success 200 {object} Education
// @Router /educations/{uuid}/ [patch]
func PartialUpdateEducationHandler(c *gin.Context) {
	uuid := c.Param("uuid")

	education, err := GetEducation(uuid)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "object with uuid=" + uuid + " not found",
			},
		)
		return
	}

	var req PartialUpdateEducationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error":  "invalid request",
				"detail": err.Error(),
			},
		)
		return
	}

	institution := fields.CharField(c, req.Institution, "institution", 100)
	major := fields.CharField(c, req.Major, "major", 100)
	degree := fields.IntChoiceField(c, req.Degree, "degree", DegreeChoices)
	startDate := fields.DateField(c, req.StartDate, "startDate")
	endDate := fields.DateField(c, req.EndDate, "endDate")

	errors := exception.GetAPIExceptions(c)

	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	education, err = UpdateEducation(
		*education,
		institution,
		degree,
		major,
		startDate,
		endDate,
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error":  "failed to update education",
				"detail": err.Error(),
			},
		)
		return
	}

	c.JSON(http.StatusOK, education)
}

// ListEducationsHandler lists all education records
// @Summary List all education records
// @Description List all education records from the database
// @Tags educations
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Param degrees query []string false "Degrees" Enums(BACHELORS,MASTERS,DOCTORATE)
// @Param ordering query []string false "Ordering" Enums(START_DATE,-START_DATE,END_DATE,-END_DATE,CREATED_AT,-CREATED_AT,DEGREE,-DEGREE)
// @Produce json
// @Success 200 {object} []Education
// @Router /educations [get]
func ListEducationsHandler(c *gin.Context) {
	limit := parameters.IntParam(c, "limit")
	offset := parameters.IntParam(c, "offset")
	degrees := parameters.IntChoiceListParam(c, "degrees", DegreeChoices)
	ordering := parameters.TextChoiceListParam(c, "ordering", EducationOrderingChoices)

	errors := exception.GetAPIExceptions(c)

	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	educations, count := ListEducation(
		limit,
		offset,
		degrees,
		ordering,
	)

	response := pagination.GetPaginatedResponse(
		c,
		limit,
		offset,
		count,
		educations,
	)

	errors = exception.GetAPIExceptions(c)
	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	c.JSON(http.StatusOK, response)
}
