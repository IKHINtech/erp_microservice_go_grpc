package service

import (
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/internal/dto"
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/internal/repository"
	pb "github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/pb/v1"
)

type DepartmentService interface {
	CreateDepartment(data *pb.CreateDepartmentRequest) (*pb.Department, error)
	UpdateDepartment(data *pb.UpdateDepartmentRequest) (*pb.Department, error)
	GetDepartment(data *pb.GetDepartmentRequest) (*pb.GetDepartmentResponse, error)
	ListDepartment(data *pb.ListDepartmentRequest) (*pb.ListDepartmentResponse, error)
	DeleteDepartment(data *pb.DeleteDepartmentRequest) error
}

type departmentServiceImpl struct {
	repo repository.DepartmentRepository
}

func NewDepartmentService(repo repository.DepartmentRepository) DepartmentService {
	return &departmentServiceImpl{
		repo: repo,
	}
}

func (s *departmentServiceImpl) CreateDepartment(department *pb.CreateDepartmentRequest) (*pb.Department, error) {
	data := dto.CreateDepartmentRequestToModel(department)
	res, err := s.repo.CreateDepartment(data)
	if err != nil {
		return nil, err
	}
	result := dto.DepartmentToProto(res)
	return result, err
}

func (s *departmentServiceImpl) GetDepartment(department *pb.GetDepartmentRequest) (*pb.GetDepartmentResponse, error) {
	res, err := s.repo.GetDepartment(department.Id)
	if err != nil {
		return nil, err
	}
	departmentProto := dto.DepartmentToProto(res)
	result := dto.GetDepartmentProtoToResponse(departmentProto)

	return result, err
}

func (s *departmentServiceImpl) ListDepartment(department *pb.ListDepartmentRequest) (*pb.ListDepartmentResponse, error) {
	res, err := s.repo.GetListDepartment()
	if err != nil {
		return nil, err
	}
	departmentsProto := make([]*pb.Department, len(res))
	for i, department := range res {
		departmentsProto[i] = dto.DepartmentToProto(department)
	}
	result := &pb.ListDepartmentResponse{
		Departments: departmentsProto,
	}
	return result, err
}

func (s *departmentServiceImpl) UpdateDepartment(department *pb.UpdateDepartmentRequest) (*pb.Department, error) {
	payload := dto.UpdateDepartmentRequestToModel(department)
	updatedDepartment, err := s.repo.UpdateDepartment(department.Id, payload)
	if err != nil {
		return nil, err
	}
	result := dto.DepartmentToProto(updatedDepartment)
	return result, nil

}

func (s *departmentServiceImpl) DeleteDepartment(department *pb.DeleteDepartmentRequest) error {
	if err := s.repo.DeleteDepartment(department.Id); err != nil {
		return err
	}
	return nil

}
