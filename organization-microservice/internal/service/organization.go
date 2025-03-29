package service

import (
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/internal/dto"
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/internal/repository"
	pb "github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/pb/v1"
)

type OrganizationService interface {
	CreateOrganization(data *pb.CreateOrganizationRequest) (*pb.Organization, error)
}

type organizationServiceImpl struct {
	repo repository.OranizationRepository
}

func NewOrganizationService(repo repository.OranizationRepository) OrganizationService {
	return &organizationServiceImpl{
		repo: repo,
	}
}

func (s *organizationServiceImpl) CreateOrganization(organization *pb.CreateOrganizationRequest) (*pb.Organization, error) {
	data := dto.CreateOrganizationRequestToModel(organization)
	res, err := s.repo.CreateOrganization(data)
	if err != nil {
		return nil, err
	}
	result := dto.OrganizationToProto(res)
	return result, err
}

func (s *organizationServiceImpl) GetOrganization(organization *pb.GetOrganizationRequest) (*pb.GetOrganizationResponse, error) {
	res, err := s.repo.GetOrganization(organization.Id)
	if err != nil {
		return nil, err
	}
	organizationProto := dto.OrganizationToProto(res)
	result := dto.GetOrganizationProtoToResponse(organizationProto)

	return result, err
}

// GetOrganization(context.Context, *GetOrganizationRequest) (*GetOrganizationResponse, error)
// ListOrganization(context.Context, *ListOrganizationRequest) (*ListOrganizationResponse, error)
// UpdateOrganization(context.Context, *UpdateOrganizationRequest) (*Organization, error)
// DeleteOrganization(context.Context, *DeleteOrganizationRequest) (*Organization, error)
