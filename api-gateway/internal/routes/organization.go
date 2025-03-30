package routes

import (
	"github.com/IKHINtech/erp_api_gateway/internal/modules/organization"
	"github.com/gin-gonic/gin"
)

func RegisterOrganizationRoutes(r *gin.Engine, organizationHandler *organization.OrganizationHandler) {
	authGroup := r.Group("/organization")
	{
		authGroup.GET("/", organizationHandler.ListOrganizations)
	}
}
