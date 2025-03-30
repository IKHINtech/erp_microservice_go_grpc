package service

import (
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/internal/dto"
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/internal/repository"
	pb "github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/pb/v1"
)

type PositionService interface {
	CreatePosition(data *pb.CreatePositionRequest) (*pb.Position, error)
	UpdatePosition(data *pb.UpdatePositionRequest) (*pb.Position, error)
	GetPosition(data *pb.GetPositionRequest) (*pb.GetPositionResponse, error)
	ListPosition(data *pb.ListPositionRequest) (*pb.ListPositionResponse, error)
	DeletePosition(data *pb.DeletePositionRequest) error
}

type positionServiceImpl struct {
	repo repository.PositionRepository
}

func NewPositionService(repo repository.PositionRepository) PositionService {
	return &positionServiceImpl{
		repo: repo,
	}
}

func (s *positionServiceImpl) CreatePosition(position *pb.CreatePositionRequest) (*pb.Position, error) {
	data := dto.CreatePositionRequestToModel(position)
	res, err := s.repo.CreatePosition(data)
	if err != nil {
		return nil, err
	}
	result := dto.PositionToProto(res)
	return result, err
}

func (s *positionServiceImpl) GetPosition(position *pb.GetPositionRequest) (*pb.GetPositionResponse, error) {
	res, err := s.repo.GetPosition(position.Id)
	if err != nil {
		return nil, err
	}
	positionProto := dto.PositionToProto(res)
	result := dto.GetPositionProtoToResponse(positionProto)

	return result, err
}

func (s *positionServiceImpl) ListPosition(position *pb.ListPositionRequest) (*pb.ListPositionResponse, error) {
	res, err := s.repo.GetListPosition()
	if err != nil {
		return nil, err
	}
	positionsProto := make([]*pb.Position, len(res))
	for i, position := range res {
		positionsProto[i] = dto.PositionToProto(position)
	}
	result := &pb.ListPositionResponse{
		Positions: positionsProto,
	}
	return result, err
}

func (s *positionServiceImpl) UpdatePosition(position *pb.UpdatePositionRequest) (*pb.Position, error) {
	payload := dto.UpdatePositionRequestToModel(position)
	updatedPosition, err := s.repo.UpdatePosition(position.Id, payload)
	if err != nil {
		return nil, err
	}
	result := dto.PositionToProto(updatedPosition)
	return result, nil

}

func (s *positionServiceImpl) DeletePosition(position *pb.DeletePositionRequest) error {
	if err := s.repo.DeletePosition(position.Id); err != nil {
		return err
	}
	return nil

}
