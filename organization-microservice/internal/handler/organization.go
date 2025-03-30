package handler

import (
	"context"

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
		UnimplementedOrganizationServiceServer: pb.UnimplementedOrganizationServiceServer{},
		organizationService:                    organizationService,
		departmentService:                      departmentService,
	}
}

func (s *OrganizationServer) CreateOrganization(ctx context.Context, data *pb.CreateOrganizationRequest) (*pb.Organization, error) {
	return s.organizationService.CreateOrganization(data)
}

func (s *OrganizationServer) ListOrganization(ctx context.Context, data *pb.ListOrganizationRequest) (*pb.ListOrganizationResponse, error) {
	return s.organizationService.ListOrganization(data)
}

func (s *OrganizationServer) GetOrganization(ctx context.Context, data *pb.GetOrganizationRequest) (*pb.GetOrganizationResponse, error) {
	return s.organizationService.GetOrganization(data)
}

func (s OrganizationServer) UpdateOrganization(ctx context.Context, data *pb.UpdateOrganizationRequest) (*pb.Organization, error) {
	return s.organizationService.UpdateOrganization(data)
}

func (s *OrganizationServer) DeleteOrganization(ctx context.Context, data *pb.DeleteOrganizationRequest) error {
	return s.organizationService.DeleteOrganization(data)
}

func (s *OrganizationServer) CreateDepartment(ctx context.Context, data *pb.CreateDepartmentRequest) (*pb.Department, error) {
	return s.departmentService.CreateDepartment(data)
}

func (s *OrganizationServer) UpdateDepartment(ctx context.Context, data *pb.UpdateDepartmentRequest) (*pb.Department, error) {
	return s.departmentService.UpdateDepartment(data)
}

func (s *OrganizationServer) GetDepartment(ctx context.Context, data *pb.GetDepartmentRequest) (*pb.GetDepartmentResponse, error) {
	return s.departmentService.GetDepartment(data)
}

func (s *OrganizationServer) ListDepartment(ctx context.Context, data *pb.ListDepartmentRequest) (*pb.ListDepartmentResponse, error) {
	return s.departmentService.ListDepartment(data)
}

func (s *OrganizationServer) DeleteDepartment(ctx context.Context, data *pb.DeleteDepartmentRequest) error {
	return s.departmentService.DeleteDepartment(data)
}
