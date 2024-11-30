package main

import (
	"log"

	_ "portfolio/docs" // Import Swagger docs

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"portfolio/config"
	"portfolio/internal/contact"
	"portfolio/internal/education"
	"portfolio/internal/experience"
	"portfolio/internal/responsibility"
	"portfolio/internal/skill"
	"portfolio/routes"
)

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Portfolio API")
// }

// @title Portfolio API
// @version 1.0
// @description API documentation for the Portfolio project
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8080
// @BasePath /api
func main() {
	log.Printf("starting server...")

	config.ConnectDatabase()

	router := gin.Default()

	// Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.SetupRoutes(router)

	config.DB.Debug().AutoMigrate(
		&contact.Contact{},
		&education.Education{},
		&experience.Experience{},
		&responsibility.Responsibility{},
		&skill.Skill{},
	)

	log.Println("Server running on http://localhost:8080")

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
