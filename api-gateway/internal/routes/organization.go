package routes

import (
	"github.com/IKHINtech/erp_api_gateway/internal/middleware"
	"github.com/IKHINtech/erp_api_gateway/internal/modules/organization"
	"github.com/gin-gonic/gin"
)

func RegisterOrganizationRoutes(r *gin.Engine, organizationHandler *organization.OrganizationHandler) {
	authGroup := r.Group("/organization")
	{
		authGroup.GET("/", organizationHandler.ListOrganizations)
	}
}

func RegisterDepartmentRoutes(r *gin.Engine, organizationHandler *organization.OrganizationHandler) {
	authGroup := r.Group("/department", middleware.AuthMiddleware())
	{
		authGroup.GET("/", organizationHandler.GetOrganization)
		authGroup.POST("/", organizationHandler.CreateDepartment)
	}
}
