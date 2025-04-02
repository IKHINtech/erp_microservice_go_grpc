package routes

import (
	"github.com/IKHINtech/erp_api_gateway/internal/middleware"
	"github.com/IKHINtech/erp_api_gateway/internal/modules/organization"
	"github.com/gin-gonic/gin"
)

func RegisterOrganizationRoutes(r *gin.Engine, organizationHandler *organization.OrganizationHandler) {
	authGroup := r.Group("/org/organization", middleware.AuthMiddleware())
	{
		authGroup.GET("/", organizationHandler.ListOrganizations)
		authGroup.POST("/", organizationHandler.CreateOrganization)
	}
}

func RegisterDepartmentRoutes(r *gin.Engine, organizationHandler *organization.OrganizationHandler) {
	authGroup := r.Group("org/department", middleware.AuthMiddleware())
	{
		authGroup.GET("/", organizationHandler.GetDepartments)
		authGroup.POST("/", organizationHandler.CreateDepartment)
	}
}
