package service

import (
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/internal/dto"
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/internal/repository"
	pb "github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/pb/v1"
)

type OrganizationService interface {
	CreateOrganization(data *pb.CreateOrganizationRequest) (*pb.Organization, error)
	UpdateOrganization(data *pb.UpdateOrganizationRequest) (*pb.Organization, error)
	GetOrganization(data *pb.GetOrganizationRequest) (*pb.GetOrganizationResponse, error)
	ListOrganization(data *pb.ListOrganizationRequest) (*pb.ListOrganizationResponse, error)
	DeleteOrganization(data *pb.DeleteOrganizationRequest) error
}

type organizationServiceImpl struct {
	repo repository.OrganizationRepository
}

func NewOrganizationService(repo repository.OrganizationRepository) OrganizationService {
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

func (s *organizationServiceImpl) ListOrganization(organization *pb.ListOrganizationRequest) (*pb.ListOrganizationResponse, error) {
	res, err := s.repo.GetListOrganization(nil)
	if err != nil {
		return nil, err
	}
	organizationsProto := make([]*pb.Organization, len(res))
	for i, organization := range res {
		organizationsProto[i] = dto.OrganizationToProto(organization)
	}
	result := &pb.ListOrganizationResponse{
		Organizations: organizationsProto,
	}
	return result, err
}

func (s *organizationServiceImpl) UpdateOrganization(organization *pb.UpdateOrganizationRequest) (*pb.Organization, error) {
	payload := dto.UpdateOrganizationRequestToModel(organization)
	updatedOrganization, err := s.repo.UpdateOrganization(organization.Id, payload)
	if err != nil {
		return nil, err
	}
	result := dto.OrganizationToProto(updatedOrganization)
	return result, nil

}

func (s *organizationServiceImpl) DeleteOrganization(organization *pb.DeleteOrganizationRequest) error {
	if err := s.repo.DeleteOrganization(organization.Id); err != nil {
		return err
	}
	return nil

}
