package server

import (
	"log"
	"net"

	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/internal/database"
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/internal/handler"
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/internal/repository"
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/internal/service"
	organizationv1 "github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/pb/v1"

	authClient "github.com/IKHINtech/erp_microservice_go_grpc/auth-microservice/client"
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/config"
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/internal/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartServer() {

	lis, err := net.Listen("tcp", "0.0.0.0:"+config.CFG.PORT) // Ganti localhost ke 0.0.0.0
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	authSdkClient, err := authClient.NewClient("0.0.0.0:"+config.CFG.AUTH_PORT, 10)
	if err != nil {
		log.Fatalf("Failed to connect to auth-microservice: %v", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.AuthInterceptor(authSdkClient)),
	)

	// di oganisazion
	repoOrganization := repository.NewOrganizationRepository(database.DB)
	serviceOrganization := service.NewOrganizationService(repoOrganization)

	// di department
	repoDepartment := repository.NewDepartmentRepository(database.DB)
	serviceDepartment := service.NewDepartmentService(repoDepartment)

	// di position
	repoPosition := repository.NewPositionRepository(database.DB)
	servicePosition := service.NewPositionService(repoPosition)

	// di work unit
	repoWorkUnit := repository.NewWorkUnitRepository(database.DB)
	serviceWorkUnit := service.NewWorkUnitService(repoWorkUnit)

	// di position hierarchy
	repoPositionHierarchy := repository.NewPositionHierarchyRepository(database.DB)
	servicePositionHierarchy := service.NewPositionHierarchyService(repoPositionHierarchy)

	organizationServer := handler.NewOrganizationServer(
		serviceOrganization,
		serviceDepartment,
		serviceWorkUnit,
		servicePositionHierarchy,
		servicePosition,
	)

	organizationv1.RegisterOrganizationServiceServer(s, &organizationServer)

	reflection.Register(s)

	log.Printf("Server running on %s", lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
