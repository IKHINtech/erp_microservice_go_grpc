package service

import (
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/internal/dto"
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/internal/repository"
	pb "github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/pb/v1"
)

type WorkUnitService interface {
	CreateWorkUnit(data *pb.CreateWorkUnitRequest) (*pb.WorkUnit, error)
	UpdateWorkUnit(data *pb.UpdateWorkUnitRequest) (*pb.WorkUnit, error)
	GetWorkUnit(data *pb.GetWorkUnitRequest) (*pb.GetWorkUnitResponse, error)
	ListWorkUnit(data *pb.ListWorkUnitRequest) (*pb.ListWorkUnitResponse, error)
	DeleteWorkUnit(data *pb.DeleteWorkUnitRequest) error
}

type WorkUnitServiceImpl struct {
	repo repository.WorkUnitRepository
}

func NewWorkUnitService(repo repository.WorkUnitRepository) WorkUnitService {
	return &WorkUnitServiceImpl{
		repo: repo,
	}
}

func (s *WorkUnitServiceImpl) CreateWorkUnit(WorkUnit *pb.CreateWorkUnitRequest) (*pb.WorkUnit, error) {
	data := dto.CreateWorkUnitRequestToModel(WorkUnit)
	res, err := s.repo.CreateWorkUnit(data)
	if err != nil {
		return nil, err
	}
	result := dto.WorkUnitToProto(res)
	return result, err
}

func (s *WorkUnitServiceImpl) GetWorkUnit(WorkUnit *pb.GetWorkUnitRequest) (*pb.GetWorkUnitResponse, error) {
	res, err := s.repo.GetWorkUnit(WorkUnit.Id)
	if err != nil {
		return nil, err
	}
	WorkUnitProto := dto.WorkUnitToProto(res)
	result := dto.GetWorkUnitProtoToResponse(WorkUnitProto)

	return result, err
}

func (s *WorkUnitServiceImpl) ListWorkUnit(WorkUnit *pb.ListWorkUnitRequest) (*pb.ListWorkUnitResponse, error) {
	res, err := s.repo.GetListWorkUnit()
	if err != nil {
		return nil, err
	}
	WorkUnitsProto := make([]*pb.WorkUnit, len(res))
	for i, WorkUnit := range res {
		WorkUnitsProto[i] = dto.WorkUnitToProto(WorkUnit)
	}
	result := &pb.ListWorkUnitResponse{
		WorkUnits: WorkUnitsProto,
	}
	return result, err
}

func (s *WorkUnitServiceImpl) UpdateWorkUnit(WorkUnit *pb.UpdateWorkUnitRequest) (*pb.WorkUnit, error) {
	payload := dto.UpdateWorkUnitRequestToModel(WorkUnit)
	updatedWorkUnit, err := s.repo.UpdateWorkUnit(WorkUnit.Id, payload)
	if err != nil {
		return nil, err
	}
	result := dto.WorkUnitToProto(updatedWorkUnit)
	return result, nil

}

func (s *WorkUnitServiceImpl) DeleteWorkUnit(WorkUnit *pb.DeleteWorkUnitRequest) error {
	if err := s.repo.DeleteWorkUnit(WorkUnit.Id); err != nil {
		return err
	}
	return nil

}
