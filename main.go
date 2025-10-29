package main

import (
	"log"

	"mcp-auth-backend-go/config"
	"mcp-auth-backend-go/models"
	"mcp-auth-backend-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// ✅ 1. Gin 모드 설정 (릴리스 모드 전환)
	gin.SetMode(gin.ReleaseMode)

	// ✅ 2. DB 연결 및 마이그레이션
	db := config.ConnectDatabase()

	if err := db.AutoMigrate(
		&models.User{},
		&models.Team{},
		&models.TeamMember{},
		&models.Project{},
		&models.Policy{},
		&models.ProjectPolicy{},
	); err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}

	log.Println("✅ Tables migrated successfully")

	r := gin.Default()

	// ✅ Users / Auth (registration, login)
	routes.RegisterUserRoutes(r)

	// ✅ Teams
	routes.RegisterTeamRoutes(r)

	// ✅ Team Members
	routes.RegisterTeamMemberRoutes(r)

	// ✅ Projects
	routes.RegisterProjectRoutes(r)

	// ✅ Policies
	routes.RegisterPolicyRoutes(r)

	// ✅ Project-Policy Relations
	routes.RegisterProjectPolicyRoutes(r)

	log.Println("🚀 Server running on http://localhost:8080")
	r.Run(":8080")
}
