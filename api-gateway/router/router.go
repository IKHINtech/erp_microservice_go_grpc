package router

import (
	"errors"
	"log"
	"net/http"

	"github.com/IKHINtech/erp_api_gateway/config"
	"github.com/IKHINtech/erp_api_gateway/internal/routes"
	"github.com/gin-gonic/gin"

	"github.com/IKHINtech/erp_api_gateway/docs"

	"github.com/IKHINtech/erp_api_gateway/internal/utils"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/IKHINtech/erp_api_gateway/internal/modules/auth"
	"github.com/IKHINtech/erp_api_gateway/internal/modules/organization"

	authClient "github.com/IKHINtech/erp_microservice_go_grpc/auth-microservice/client"
	orgClient "github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/client"
)

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
	// Register routes here
	authClient, err := authClient.NewClient("0.0.0.0:"+cfg.AUTH_PORT, 10)
	if err != nil {
		log.Fatalf("Failed to connect to auth service: %v", err)
	}

	organizationClient, err := orgClient.NewOrganizationClient("0.0.0.0:"+cfg.ORGANIZATION_PORT, 10)
	if err != nil {
		log.Fatalf("Failed to connect to auth service: %v", err)
	}
	authHandler := auth.NewHandler(authClient)
	organizationHandler := organization.NewOrganizationHandler(organizationClient)

	routes.RegisterAuthRoutes(r, authHandler)
	routes.RegisterOrganizationRoutes(r, organizationHandler)
	routes.RegisterDepartmentRoutes(r, organizationHandler)

	r.GET("/", func(ctx *gin.Context) {
		utils.RespondSuccess(ctx, http.StatusOK, "API Gateway", nil)
	})
	r.NoRoute(func(c *gin.Context) {
		utils.RespondError(
			c, http.StatusNotFound, "NOT_FOUND",
			errors.New("route not found"),
			nil,
		)
	})

	ginSwagger.WrapHandler(swaggerfiles.Handler)
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
