package handler

import (
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/internal/service"
	pb "github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/pb/v1"
)

type OrganizationServer struct {
	pb.UnimplementedOrganizationServiceServer
	organizationService service.OrganizationService
	departmentService   service.DepartmentService
}

func NewOrganizationServer(
	organizationService service.OrganizationService,
	departmentService service.DepartmentService,
) OrganizationServer {
	return OrganizationServer{
		organizationService: organizationService,
		departmentService:   departmentService,
	}
}
