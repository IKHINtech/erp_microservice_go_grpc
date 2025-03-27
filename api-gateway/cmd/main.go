package main

import (
	"log"

	config "github.com/IKHINtech/erp_api_gateway/config"
	"github.com/IKHINtech/erp_api_gateway/docs"
	"github.com/IKHINtech/erp_api_gateway/internal/middleware"
	"github.com/IKHINtech/erp_api_gateway/internal/modules/auth"
	"github.com/IKHINtech/erp_api_gateway/internal/routes"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			 API Gateway Documentation
//	@version		1.0
//	@description	This is a API documentation for ERP microservice.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	sarikhin@yahoo.co.id

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host						localhost:8080
// @BasePath					/
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Gagal load config:", err)
	}
	authClient, err := auth.NewAuthClient("0.0.0.0:" + cfg.AUTH_PORT)
	if err != nil {
		log.Fatalf("Failed to connect to auth service: %v", err)
	}

	r := gin.Default()

	r.Use(middleware.ErrorHandler())

	authHandler := auth.NewHandler(authClient)

	routes.RegisterAuthRoutes(r, authHandler)

	r.GET("/", func(ctx *gin.Context) {
		ctx.Set("response", "API GATEWAY")
	})
	r.NoRoute(func(c *gin.Context) {
		c.Set("code", 404)
		c.Set("response", "Route Not Found")
	})

	ginSwagger.WrapHandler(swaggerfiles.Handler)
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	if err := r.Run(":" + cfg.GATEWAY_PORT); err != nil {
		log.Fatal(err)
	}
}
