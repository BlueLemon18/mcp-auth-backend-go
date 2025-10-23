package main

import (
	"log"

	"mcp-auth-backend-go/config"
	"mcp-auth-backend-go/models"
)

func main() {
	db := config.ConnectDatabase() // 네가 이미 만든 함수명 유지

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
}
