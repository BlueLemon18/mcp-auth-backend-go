package routes

import (
	"mcp-auth-backend-go/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
	}
}
