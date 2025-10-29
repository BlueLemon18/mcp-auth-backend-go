package handlers

import (
	"mcp-auth-backend-go/config"
	"mcp-auth-backend-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 팀에 사용자 추가
func AddTeamMember(c *gin.Context) {
	teamID := c.Param("id")
	var member models.TeamMember

	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	member.TeamID = uint64(parseUint(teamID)) // ✅ uint64로 변환
	db := config.ConnectDatabase()
	if err := db.Create(&member).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add member"})
		return
	}

	c.JSON(http.StatusCreated, member)
}

// 팀원 목록 조회
func GetTeamMembers(c *gin.Context) {
	teamID := c.Param("id")
	db := config.ConnectDatabase()

	var members []models.TeamMember
	db.Where("team_id = ?", teamID).Find(&members)
	c.JSON(http.StatusOK, members)
}

// 팀원 역할 변경
func UpdateTeamMemberRole(c *gin.Context) {
	teamID := c.Param("id")
	userID := c.Param("user_id")
	db := config.ConnectDatabase()

	var member models.TeamMember
	if err := db.Where("team_id = ? AND user_id = ?", teamID, userID).First(&member).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "member not found"})
		return
	}

	var body struct {
		Role string `json:"role"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	member.Role = body.Role
	db.Save(&member)
	c.JSON(http.StatusOK, member)
}

// 팀원 제거
func RemoveTeamMember(c *gin.Context) {
	teamID := c.Param("id")
	userID := c.Param("user_id")
	db := config.ConnectDatabase()

	if err := db.Where("team_id = ? AND user_id = ?", teamID, userID).
		Delete(&models.TeamMember{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to remove member"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "member removed"})
}

// 도우미 함수: 문자열을 uint로 변환
func parseUint(s string) uint {
	var v uint
	for _, c := range s {
		v = v*10 + uint(c-'0')
	}
	return v
}
