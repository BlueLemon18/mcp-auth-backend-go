package handlers

import (
	"mcp-auth-backend-go/config"
	"mcp-auth-backend-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTeam(c *gin.Context) {
	var team models.Team
	if err := c.ShouldBindJSON(&team); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := config.ConnectDatabase()
	if err := db.Create(&team).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create team"})
		return
	}
	c.JSON(http.StatusCreated, team)
}

func GetAllTeams(c *gin.Context) {
	db := config.ConnectDatabase()
	var teams []models.Team
	db.Find(&teams)
	c.JSON(http.StatusOK, teams)
}

func GetTeamByID(c *gin.Context) {
	id := c.Param("id")
	db := config.ConnectDatabase()
	var team models.Team
	if err := db.First(&team, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "team not found"})
		return
	}
	c.JSON(http.StatusOK, team)
}

func UpdateTeam(c *gin.Context) {
	id := c.Param("id")
	var team models.Team
	db := config.ConnectDatabase()
	if err := db.First(&team, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "team not found"})
		return
	}
	if err := c.ShouldBindJSON(&team); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&team)
	c.JSON(http.StatusOK, team)
}

func DeleteTeam(c *gin.Context) {
	id := c.Param("id")
	db := config.ConnectDatabase()
	if err := db.Delete(&models.Team{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete team"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "team deleted"})
}
