package main

import (
	"log"

	config "github.com/IKHINtech/erp_api_gateway/config"
	"github.com/IKHINtech/erp_api_gateway/internal/middleware"
	"github.com/IKHINtech/erp_api_gateway/internal/modules/auth"
	"github.com/IKHINtech/erp_api_gateway/internal/routes"
	"github.com/gin-gonic/gin"
)

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

	if err := r.Run(":" + cfg.GATEWAY_PORT); err != nil {
		log.Fatal(err)
	}
}
