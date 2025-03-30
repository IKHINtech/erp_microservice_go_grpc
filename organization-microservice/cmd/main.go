package main

import (
	"log"
	"net"

	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/config"
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/internal/database"
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/internal/handler"
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/internal/repository"
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/internal/service"
	organizationv1 "github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/pb/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Gagal load config:", err)
	}

	// connect to database
	// 2025-03-30
	// 2025-03-30 08:36
	database.Connect()

	// migrasi database
	database.Migrate()

	lis, err := net.Listen("tcp", "0.0.0.0:"+cfg.PORT) // Ganti localhost ke 0.0.0.0
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	// di oganisazion
	repoOrganization := repository.NewOrganizationRepository(database.DB)
	serviceOrganization := service.NewOrganizationService(repoOrganization)

	// di department
	repoDepartment := repository.NewDepartmentRepository(database.DB)
	serviceDepartment := service.NewDepartmentService(repoDepartment)

	organizationv1.RegisterOrganizationServiceServer(s, handler.NewOrganizationServer(serviceOrganization, serviceDepartment))

	reflection.Register(s)

	log.Printf("Server running on %s", lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
