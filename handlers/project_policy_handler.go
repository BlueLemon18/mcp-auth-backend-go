package handlers

import (
	"net/http"
	"strconv"

	"mcp-auth-backend-go/config"
	"mcp-auth-backend-go/models"

	"github.com/gin-gonic/gin"
)

// 프로젝트에 정책 연결
func LinkPolicyToProject(c *gin.Context) {
	projectIDStr := c.Param("project_id")

	// 문자열 → uint 변환
	projectID64, err := strconv.ParseUint(projectIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid project_id"})
		return
	}
	projectID := uint(projectID64)

	// body 파싱
	var body struct {
		PolicyID uint `json:"policy_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	// 명시적 타입 지정
	link := models.ProjectPolicy{
		ProjectID: uint64(projectID),
		PolicyID:  uint64(body.PolicyID),
		IsActive:  true,
	}

	db := config.ConnectDatabase()
	if err := db.Create(&link).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to link policy"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "policy linked successfully",
		"link":    link,
	})
}

// 특정 프로젝트에 연결된 정책 조회
func GetLinkedPolicies(c *gin.Context) {
	projectIDStr := c.Param("project_id")
	db := config.ConnectDatabase()

	var linked []models.ProjectPolicy
	if err := db.Where("project_id = ?", projectIDStr).Find(&linked).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch linked policies"})
		return
	}

	c.JSON(http.StatusOK, linked)
}

// 연결 상태(is_active) 수정 (토글)
func UpdatePolicyLinkStatus(c *gin.Context) {
	projectIDStr := c.Param("project_id")
	policyIDStr := c.Param("policy_id")

	db := config.ConnectDatabase()

	var link models.ProjectPolicy
	if err := db.Where("project_id = ? AND policy_id = ?", projectIDStr, policyIDStr).First(&link).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "link not found"})
		return
	}

	link.IsActive = !link.IsActive
	if err := db.Save(&link).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update link status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "link status updated",
		"link":    link,
	})
}

// 프로젝트-정책 연결 해제
func UnlinkPolicyFromProject(c *gin.Context) {
	projectIDStr := c.Param("project_id")
	policyIDStr := c.Param("policy_id")

	db := config.ConnectDatabase()
	if err := db.Where("project_id = ? AND policy_id = ?", projectIDStr, policyIDStr).
		Delete(&models.ProjectPolicy{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to unlink policy"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "policy unlinked successfully"})
}
