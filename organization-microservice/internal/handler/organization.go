package handler

import (
	"context"

	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/internal/service"
	pb "github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/pb/v1"
)

type OrganizationServer struct {
	pb.UnimplementedOrganizationServiceServer
	organizationService      service.OrganizationService
	departmentService        service.DepartmentService
	workUnitService          service.WorkUnitService
	positionHierarchyService service.PositionHierarchyService
	positionService          service.PositionService
}

func NewOrganizationServer(
	organizationService service.OrganizationService,
	departmentService service.DepartmentService,
	workUnitService service.WorkUnitService,
	positionHierarchyService service.PositionHierarchyService,
	positionService service.PositionService,
) OrganizationServer {
	return OrganizationServer{
		UnimplementedOrganizationServiceServer: pb.UnimplementedOrganizationServiceServer{},
		organizationService:                    organizationService,
		departmentService:                      departmentService,
		workUnitService:                        workUnitService,
		positionHierarchyService:               positionHierarchyService,
		positionService:                        positionService,
	}
}

// // Oraganization Endpoint

// CreateOrganization(context.Context, *CreateOrganizationRequest) (*Organization, error)
func (s *OrganizationServer) CreateOrganization(ctx context.Context, data *pb.CreateOrganizationRequest) (*pb.Organization, error) {
	return s.organizationService.CreateOrganization(data)
}

// ListOrganization(context.Context, *ListOrganizationRequest) (*ListOrganizationResponse, error)
func (s *OrganizationServer) ListOrganization(ctx context.Context, data *pb.ListOrganizationRequest) (*pb.ListOrganizationResponse, error) {
	return s.organizationService.ListOrganization(data)
}

// GetOrganization(context.Context, *GetOrganizationRequest) (*GetOrganizationResponse, error)
func (s *OrganizationServer) GetOrganization(ctx context.Context, data *pb.GetOrganizationRequest) (*pb.GetOrganizationResponse, error) {
	return s.organizationService.GetOrganization(data)
}

// UpdateOrganization(context.Context, *UpdateOrganizationRequest) (*Organization, error)
func (s OrganizationServer) UpdateOrganization(ctx context.Context, data *pb.UpdateOrganizationRequest) (*pb.Organization, error) {
	return s.organizationService.UpdateOrganization(data)
}

// DeleteOrganization(context.Context, *DeleteOrganizationRequest) (*Organization, error)
func (s *OrganizationServer) DeleteOrganization(ctx context.Context, data *pb.DeleteOrganizationRequest) (*pb.Organization, error) {
	return nil, s.organizationService.DeleteOrganization(data)
}

// 	// Department Endpoint

// CreateDepartment(context.Context, *CreateDepartmentRequest) (*Department, error)
func (s *OrganizationServer) CreateDepartment(ctx context.Context, data *pb.CreateDepartmentRequest) (*pb.Department, error) {
	return s.departmentService.CreateDepartment(data)
}

// UpdateDepartment(context.Context, *UpdateDepartmentRequest) (*Department, error)
func (s *OrganizationServer) UpdateDepartment(ctx context.Context, data *pb.UpdateDepartmentRequest) (*pb.Department, error) {
	return s.departmentService.UpdateDepartment(data)
}

// GetDepartment(context.Context, *GetDepartmentRequest) (*GetDepartmentResponse, error)
func (s *OrganizationServer) GetDepartment(ctx context.Context, data *pb.GetDepartmentRequest) (*pb.GetDepartmentResponse, error) {
	return s.departmentService.GetDepartment(data)
}

// ListDepartment(context.Context, *ListDepartmentRequest) (*ListDepartmentResponse, error)
func (s *OrganizationServer) ListDepartment(ctx context.Context, data *pb.ListDepartmentRequest) (*pb.ListDepartmentResponse, error) {
	return s.departmentService.ListDepartment(data)
}

// DeleteDepartment(context.Context, *DeleteDepartmentRequest) (*Department, error)
func (s *OrganizationServer) DeleteDepartment(ctx context.Context, data *pb.DeleteDepartmentRequest) (*pb.Department, error) {
	return nil, s.departmentService.DeleteDepartment(data)
}

// // WorkUnit Endpoint

// CreateWorkUnit(context.Context, *CreateWorkUnitRequest) (*WorkUnit, error)
func (s *OrganizationServer) CreateWorkUnit(ctx context.Context, data *pb.CreateWorkUnitRequest) (*pb.WorkUnit, error) {
	return s.workUnitService.CreateWorkUnit(data)
}

// GetWorkUnit(context.Context, *GetWorkUnitRequest) (*GetWorkUnitResponse, error)
func (s *OrganizationServer) GetWorkUnit(ctx context.Context, data *pb.GetWorkUnitRequest) (*pb.GetWorkUnitResponse, error) {
	return s.workUnitService.GetWorkUnit(data)
}

// ListWorkUnit(context.Context, *ListWorkUnitRequest) (*ListWorkUnitResponse, error)
func (s *OrganizationServer) ListWorkUnit(ctx context.Context, data *pb.ListWorkUnitRequest) (*pb.ListWorkUnitResponse, error) {
	return s.workUnitService.ListWorkUnit(data)
}

// DeleteWorkUnit(context.Context, *DeleteWorkUnitRequest) (*WorkUnit, error)
func (s *OrganizationServer) DeleteWorkUnit(ctx context.Context, data *pb.DeleteWorkUnitRequest) (*pb.WorkUnit, error) {
	return nil, s.workUnitService.DeleteWorkUnit(data)
}

// // Position Endpoint
//
// CreatePosition(context.Context, *CreatePositionRequest) (*Position, error)
func (s *OrganizationServer) CreatePosition(ctx context.Context, data *pb.CreatePositionRequest) (*pb.Position, error) {
	return s.positionService.CreatePosition(data)
}

// GetPosition(context.Context, *GetPositionRequest) (*GetPositionResponse, error)
func (s *OrganizationServer) GetPosition(ctx context.Context, data *pb.GetPositionRequest) (*pb.GetPositionResponse, error) {
	return s.positionService.GetPosition(data)
}

// ListPosition(context.Context, *ListPositionRequest) (*ListPositionResponse, error)
func (s *OrganizationServer) ListPosition(ctx context.Context, data *pb.ListPositionRequest) (*pb.ListPositionResponse, error) {
	return s.positionService.ListPosition(data)
}

// UpdatePosition(context.Context, *UpdatePositionRequest) (*Position, error)
func (s *OrganizationServer) UpdatePosition(ctx context.Context, data *pb.UpdatePositionRequest) (*pb.Position, error) {
	return s.positionService.UpdatePosition(data)
}

// DeletePosition(context.Context, *DeletePositionRequest) (*Position, error)
func (s *OrganizationServer) DeletePosition(ctx context.Context, data *pb.DeletePositionRequest) (*pb.Position, error) {
	return nil, s.positionService.DeletePosition(data)
}

// // Work Hierarchy Endpoint
//
// CreatePositionHierarchy(context.Context, *CreatePositionHierarchyRequest) (*PositionHierarchy, error)
func (s *OrganizationServer) CreatePositionHierarchy(ctx context.Context, data *pb.CreatePositionHierarchyRequest) (*pb.PositionHierarchy, error) {
	return s.positionHierarchyService.CreatePositionHierarchy(data)
}

// GetPositionHierarchyBySuperior(context.Context, *GetPositionHierarchyBySuperiorRequest) (*ListPositionHierarchyResponse, error)
func (s *OrganizationServer) GetPositionHierarchyBySuperior(ctx context.Context, data *pb.GetPositionHierarchyBySuperiorRequest) (*pb.ListPositionHierarchyResponse, error) {
	return s.positionHierarchyService.GetPositionHierarchyBySuperior(data)
}

// GetPositionHierarchyBySubordinated(context.Context, *GetPositionHierarchyBySubordinatedRequest) (*ListPositionHierarchyResponse, error)
func (s *OrganizationServer) GetPositionHierarchyBySubordinated(ctx context.Context, data *pb.GetPositionHierarchyBySubordinatedRequest) (*pb.ListPositionHierarchyResponse, error) {
	return s.positionHierarchyService.GetPositionHierarchyBySubordinated(data)
}

// DeletePositionHierarchy(context.Context, *DeletePositionHierarchyRequest) (*PositionHierarchy, error)
func (s *OrganizationServer) DeletePositionHierarchy(ctx context.Context, data *pb.DeletePositionHierarchyRequest) (*pb.PositionHierarchy, error) {
	return nil, s.positionHierarchyService.DeletePositionHierarchy(data)
}
