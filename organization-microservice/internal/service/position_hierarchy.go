package service

import (
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/internal/dto"
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/internal/repository"
	pb "github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/pb/v1"
)

type PositionHierarchyService interface {
	CreatePositionHierarchy(data *pb.CreatePositionHierarchyRequest) (*pb.PositionHierarchy, error)
	GetPositionHierarchyBySuperior(data *pb.GetPositionHierarchyBySuperiorRequest) (*pb.ListPositionHierarchyResponse, error)
	GetPositionHierarchyBySubordinated(data *pb.GetPositionHierarchyBySubordinatedRequest) (*pb.ListPositionHierarchyResponse, error)
	DeletePositionHierarchy(data *pb.DeletePositionHierarchyRequest) error
}

type positionHierarchyServiceImpl struct {
	repo repository.PositionHierarchyRepository
}

func NewPositionHierarchyService(repo repository.PositionHierarchyRepository) PositionHierarchyService {
	return &positionHierarchyServiceImpl{
		repo: repo,
	}
}

func (s *positionHierarchyServiceImpl) CreatePositionHierarchy(PositionHierarchy *pb.CreatePositionHierarchyRequest) (*pb.PositionHierarchy, error) {
	data := dto.CreatePositionHierarchyRequestToModel(PositionHierarchy)
	res, err := s.repo.CreatePositionHierarchy(data)
	if err != nil {
		return nil, err
	}
	result := dto.PositionHierarchyToProto(res)
	return result, err
}

func (s *positionHierarchyServiceImpl) GetPositionHierarchyBySuperior(data *pb.GetPositionHierarchyBySuperiorRequest) (*pb.ListPositionHierarchyResponse, error) {
	result, err := s.repo.GetListPositionHierarchy(&data.SuperiorPositionId, nil)
	if err != nil {
		return nil, err
	}

	resultProto := make([]*pb.PositionHierarchy, len(result))
	for i, result := range result {
		resultProto[i] = dto.PositionHierarchyToProto(result)
	}

	return &pb.ListPositionHierarchyResponse{PositionHierarchy: resultProto}, err
}

func (s *positionHierarchyServiceImpl) GetPositionHierarchyBySubordinated(data *pb.GetPositionHierarchyBySubordinatedRequest) (*pb.ListPositionHierarchyResponse, error) {

	result, err := s.repo.GetListPositionHierarchy(nil, &data.SubordinatedPositionId)
	if err != nil {
		return nil, err
	}

	resultProto := make([]*pb.PositionHierarchy, len(result))
	for i, result := range result {
		resultProto[i] = dto.PositionHierarchyToProto(result)
	}

	return &pb.ListPositionHierarchyResponse{PositionHierarchy: resultProto}, err
}

func (s *positionHierarchyServiceImpl) DeletePositionHierarchy(PositionHierarchy *pb.DeletePositionHierarchyRequest) error {
	if err := s.repo.DeletePositionHierarchy(PositionHierarchy.Id); err != nil {
		return err
	}
	return nil

}
