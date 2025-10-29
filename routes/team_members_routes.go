package routes

import (
	"mcp-auth-backend-go/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterTeamMemberRoutes(r *gin.Engine) {
	members := r.Group("/teams/:id/members")
	{
		members.POST("", handlers.AddTeamMember)
		members.GET("", handlers.GetTeamMembers)
		members.PUT("/:user_id", handlers.UpdateTeamMemberRole)
		members.DELETE("/:user_id", handlers.RemoveTeamMember)
	}
}
