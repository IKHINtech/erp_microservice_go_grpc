package main

import (
	"log"

	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/config"
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/internal/database"
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/server"
)

func main() {
	_, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Gagal load config:", err)
	}
	// connect to database
	database.Connect()
	// migrasi database
	database.Migrate()

	// start server
	server.StartServer()
}
