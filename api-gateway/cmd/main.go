package main

import (
	"log"

	"github.com/IKHINtech/erp_api_gateway/internal/modules/auth"
	"github.com/IKHINtech/erp_api_gateway/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Init Auth Client
	authClient, err := auth.NewAuthClient("0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to connect to auth service: %v", err)
	}

	// 2. Create Handlers
	authHandler := auth.NewHandler(authClient)

	// 3. Setup Gin
	r := gin.Default()

	// 4. Register Routes
	routes.RegisterAuthRoutes(r, authHandler)

	// 5. Start Server
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
