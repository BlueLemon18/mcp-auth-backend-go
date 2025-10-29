package routes

import (
	"mcp-auth-backend-go/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterProjectRoutes(r *gin.Engine) {
	projects := r.Group("/projects")
	{
		projects.POST("", handlers.CreateProject)
		projects.GET("", handlers.GetAllProjects)
		projects.GET("/:id", handlers.GetProjectByID)
		projects.PUT("/:id", handlers.UpdateProject)
		projects.DELETE("/:id", handlers.DeleteProject)
	}
}
