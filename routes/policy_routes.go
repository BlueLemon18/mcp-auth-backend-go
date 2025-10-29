package routes

import (
	"mcp-auth-backend-go/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterPolicyRoutes(r *gin.Engine) {
	policies := r.Group("/policies")
	{
		policies.POST("", handlers.CreatePolicy)
		policies.GET("", handlers.GetAllPolicies)
		policies.GET("/:id", handlers.GetPolicyByID)
		policies.PUT("/:id", handlers.UpdatePolicy)
		policies.DELETE("/:id", handlers.DeletePolicy)
	}
}
