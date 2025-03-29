package service

import (
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/models"
	pb "github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/pb/v1"
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/repository"
)

type OrganizationService struct {
	pb.UnimplementedOrganizationServiceServer
	repo repository.OranizationRepository
}

func NewOrganizationService(repo repository.OranizationRepository) *OrganizationService {
	return &OrganizationService{
		repo: repo,
	}
}

func (s *OrganizationService) CreateOrganization(organization *pb.CreateOrganizationRequest) (*pb.GetOrganizationResponse, error) {
	data := models.Organization{
		Name:                 organization.Name,
		Description:          organization.Description,
		Type:                 organization.Type,
		ParentOrganizationID: &organization.ParentOrganizationId,
	}
	err := s.repo.CreateOrganization(&data)
	return &pb.GetOrganizationResponse{}, err
}
