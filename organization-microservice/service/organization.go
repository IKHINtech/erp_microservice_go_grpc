package service

import (
	"context"

	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/models"
	pb "github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/pb/v1"
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/repository"
)

// CreateOrganization(context.Context, *CreateOrganizationRequest) (*Organization, error)
// // Department Endpoint
// CreateDepartment(context.Context, *CreateDepartmentRequest) (*Department, error)
// GetDepartment(context.Context, *GetDepartmentRequest) (*GetDepartmentResponse, error)
// ListDepartment(context.Context, *ListDepartmentRequest) (*ListDepartmentResponse, error)
// UpdateDepartment(context.Context, *UpdateDepartmentRequest) (*Department, error)
// DeleteDepartment(context.Context, *DeleteDepartmentRequest) (*Department, error)
// // WorkUnit Endpoint
// CreateWorkUnit(context.Context, *CreateWorkUnitRequest) (*WorkUnit, error)
// GetWorkUnit(context.Context, *GetWorkUnitRequest) (*GetWorkUnitResponse, error)
// ListWorkUnit(context.Context, *ListWorkUnitRequest) (*ListWorkUnitResponse, error)
// DeleteWorkUnit(context.Context, *DeleteWorkUnitRequest) (*WorkUnit, error)
// // Position Endpoint
// CreatePosition(context.Context, *CreatePositionRequest) (*Position, error)
// GetPosition(context.Context, *GetPositionRequest) (*GetPositionResponse, error)
// ListPosition(context.Context, *ListPositionRequest) (*ListPositionResponse, error)
// UpdatePosition(context.Context, *UpdatePositionRequest) (*Position, error)
// DeletePosition(context.Context, *DeletePositionRequest) (*Position, error)
// // Work Hierarchy Endpoint
// CreatePositionHierarchy(context.Context, *CreatePositionHierarchyRequest) (*PositionHierarchy, error)
// GetPositionHierarchyBySuperior(context.Context, *GetPositionHierarchyBySuperiorRequest) (*ListPositionHierarchyResponse, error)
// GetPositionHierarchyBySubordinated(context.Context, *GetPositionHierarchyBySubordinatedRequest) (*ListPositionHierarchyResponse, error)
// DeletePositionHierarchy(context.Context, *DeletePositionHierarchyRequest) (*PositionHierarchy, error)

type OrganizationService struct {
	pb.UnimplementedOrganizationServiceServer
	repo repository.OranizationRepository
}

func NewOrganizationService(repo repository.OranizationRepository) *OrganizationService {
	return &OrganizationService{
		UnimplementedOrganizationServiceServer: pb.UnimplementedOrganizationServiceServer{},
		repo:                                   repo,
	}
}

func (s *OrganizationService) CreateOrganization(context context.Context, organization *pb.CreateOrganizationRequest) (*pb.Organization, error) {
	data := models.Organization{
		Name:                 organization.Name,
		Description:          organization.Description,
		Type:                 organization.Type,
		ParentOrganizationID: &organization.ParentOrganizationId,
	}
	res, err := s.repo.CreateOrganization(&data)
	if err != nil {
		return nil, err
	}

	result := pb.Organization{
		Id: res.ID,
	}
	return &result, err
}

// GetOrganization(context.Context, *GetOrganizationRequest) (*GetOrganizationResponse, error)
// ListOrganization(context.Context, *ListOrganizationRequest) (*ListOrganizationResponse, error)
// UpdateOrganization(context.Context, *UpdateOrganizationRequest) (*Organization, error)
// DeleteOrganization(context.Context, *DeleteOrganizationRequest) (*Organization, error)
func (s *OrganizationService) GetOrganization(context context.Context, organization *pb.GetOrganizationRequest) (*pb.GetOrganizationResponse, error) {

	return nil, nil
}
