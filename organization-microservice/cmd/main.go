package main

import (
	"log"
	"net"

	"github.com/IKHINtech/erp_microservice_go_grpc/auth-microservice/internal/repositories"
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/config"
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/database"
	organizationv1 "github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/pb/v1"
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Gagal load config:", err)
	}

	database.Connect()
	database.Migrate()

	lis, err := net.Listen("tcp", "0.0.0.0:"+cfg.PORT) // Ganti localhost ke 0.0.0.0
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	organizationRepo := repositories.NewOrganizationRepository(database.DB)
	oranizationService := service.NewOrganizationService(organizationRepo)

	organizationv1.RegisterOrganizationServiceServer(s, oranizationService)

	reflection.Register(s)

	log.Printf("Server running on %s", lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
