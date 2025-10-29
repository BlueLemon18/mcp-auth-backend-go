package routes

import (
	"mcp-auth-backend-go/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterProjectPolicyRoutes(r *gin.Engine) {
	pp := r.Group("/projects/:id/policies")
	{
		pp.POST("", handlers.LinkPolicyToProject)
		pp.GET("", handlers.GetLinkedPolicies)
		pp.PUT("/:policy_id", handlers.UpdatePolicyLinkStatus)
		pp.DELETE("/:policy_id", handlers.UnlinkPolicyFromProject)
	}
}
