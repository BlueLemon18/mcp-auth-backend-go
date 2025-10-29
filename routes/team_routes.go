package routes

import (
	"mcp-auth-backend-go/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterTeamRoutes(r *gin.Engine) {
	teams := r.Group("/teams")
	{
		teams.POST("", handlers.CreateTeam)
		teams.GET("", handlers.GetAllTeams)
		teams.GET("/:id", handlers.GetTeamByID)
		teams.PUT("/:id", handlers.UpdateTeam)
		teams.DELETE("/:id", handlers.DeleteTeam)
	}
}
