package routes

import (
	"github.com/IKHINtech/erp_api_gateway/internal/modules/auth"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine, authHandler *auth.Handler) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", authHandler.Login)
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/validate", authHandler.ValidateToken)
	}
}
