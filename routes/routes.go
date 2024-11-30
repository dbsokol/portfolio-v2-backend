package routes

import (
	"portfolio/internal/contact"
	"portfolio/internal/education"
	"portfolio/internal/experience"
	"portfolio/internal/skill"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")

	api.GET("/contacts", contact.ListContactsHandler)
	api.POST("/contacts", contact.CreateContactHandler)

	api.GET("/educations", education.ListEducationsHandler)
	api.POST("/educations", education.CreateEducationHandler)
	api.DELETE("/educations/:uuid/", education.DeleteEducationHandler)
	api.PATCH("/educations/:uuid/", education.PartialUpdateEducationHandler)

	api.GET("/experiences", experience.ListExperiencesHandler)

	api.GET("/skills", skill.ListSkillsHandler)
	api.POST("/skills", skill.CreateSkillHandler)
	api.DELETE("/skills/:uuid/", skill.DeleteSkillHandler)
	api.PATCH("/skills/:uuid/", skill.PartialUpdateSkillHandler)

}
