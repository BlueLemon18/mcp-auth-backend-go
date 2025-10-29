package handlers

import (
	"mcp-auth-backend-go/config"
	"mcp-auth-backend-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 정책 생성
func CreatePolicy(c *gin.Context) {
	var policy models.Policy
	if err := c.ShouldBindJSON(&policy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := config.ConnectDatabase()
	if err := db.Create(&policy).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create policy"})
		return
	}
	c.JSON(http.StatusCreated, policy)
}

// 전체 정책 조회
func GetAllPolicies(c *gin.Context) {
	db := config.ConnectDatabase()
	var policies []models.Policy
	db.Find(&policies)
	c.JSON(http.StatusOK, policies)
}

// 단일 정책 조회
func GetPolicyByID(c *gin.Context) {
	id := c.Param("id")
	db := config.ConnectDatabase()
	var policy models.Policy
	if err := db.First(&policy, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "policy not found"})
		return
	}
	c.JSON(http.StatusOK, policy)
}

// 정책 수정
func UpdatePolicy(c *gin.Context) {
	id := c.Param("id")
	db := config.ConnectDatabase()
	var policy models.Policy
	if err := db.First(&policy, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "policy not found"})
		return
	}
	if err := c.ShouldBindJSON(&policy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&policy)
	c.JSON(http.StatusOK, policy)
}

// 정책 삭제
func DeletePolicy(c *gin.Context) {
	id := c.Param("id")
	db := config.ConnectDatabase()
	if err := db.Delete(&models.Policy{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete policy"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "policy deleted"})
}
