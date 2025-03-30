package dto

import (
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/internal/models"
	pb "github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/pb/v1"
)

func DepartmentToProto(data *models.Department) *pb.Department {
	return &pb.Department{
		Id:          data.ID,
		Name:        data.Name,
		Description: data.Description,
	}
}

func GetDepartmentProtoToResponse(data *pb.Department) *pb.GetDepartmentResponse {
	return &pb.GetDepartmentResponse{
		Department: data,
	}
}

func CreateDepartmentRequestToModel(data *pb.CreateDepartmentRequest) *models.Department {
	return &models.Department{
		Name:        data.Name,
		Description: data.Description,
	}
}

func UpdateDepartmentRequestToModel(data *pb.UpdateDepartmentRequest) *models.Department {
	res := models.Department{
		Name:        data.Name,
		Description: data.Description,
	}
	res.ID = data.Id
	return &res
}
