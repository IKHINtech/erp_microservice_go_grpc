package main

import (
	"log"

	config "github.com/IKHINtech/erp_api_gateway/config"
	"github.com/IKHINtech/erp_api_gateway/internal/middleware"
	"github.com/IKHINtech/erp_api_gateway/router"
	"github.com/gin-gonic/gin"
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
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Gagal load config:", err)
	}

	r := gin.Default()

	r.Use(middleware.ErrorHandler())

	router.RegisterRoutes(r, &cfg)

	if err := r.Run(":" + cfg.GATEWAY_PORT); err != nil {
		log.Fatal(err)
	}

}
