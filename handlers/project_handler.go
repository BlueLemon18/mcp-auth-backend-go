package handlers

import (
	"mcp-auth-backend-go/config"
	"mcp-auth-backend-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProject(c *gin.Context) {
	var p models.Project
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := config.ConnectDatabase()
	db.Create(&p)
	c.JSON(http.StatusCreated, p)
}

func GetAllProjects(c *gin.Context) {
	db := config.ConnectDatabase()
	var projects []models.Project
	db.Find(&projects)
	c.JSON(http.StatusOK, projects)
}

func GetProjectByID(c *gin.Context) {
	id := c.Param("id")
	db := config.ConnectDatabase()
	var project models.Project
	if err := db.First(&project, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "project not found"})
		return
	}
	c.JSON(http.StatusOK, project)
}

func UpdateProject(c *gin.Context) {
	id := c.Param("id")
	db := config.ConnectDatabase()
	var project models.Project
	if err := db.First(&project, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "project not found"})
		return
	}
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&project)
	c.JSON(http.StatusOK, project)
}

func DeleteProject(c *gin.Context) {
	id := c.Param("id")
	db := config.ConnectDatabase()
	if err := db.Delete(&models.Project{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete project"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "project deleted"})
}
