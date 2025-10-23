package main

import (
	"log"

	"mcp-auth-backend-go/config"
	"mcp-auth-backend-go/models"
	"mcp-auth-backend-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDatabase()

	if err := db.AutoMigrate(
		&models.User{},
		&models.Team{},
		&models.Policy{},
		&models.Project{},
		&models.TeamMember{},
		&models.ProjectPolicy{},
	); err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}

	log.Println("✅ Tables migrated successfully")

	r := gin.Default()
	routes.RegisterAuthRoutes(r)

	log.Println("🚀 Server running on http://localhost:8080")
	r.Run(":8080")
}
