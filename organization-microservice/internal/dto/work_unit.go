package dto

import (
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/internal/models"
	pb "github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/pb/v1"
)

func WorkUnitToProto(data *models.WorkUnit) *pb.WorkUnit {
	return &pb.WorkUnit{
		Id:           data.ID,
		Name:         data.Name,
		DepartmentId: data.DepartmentID,
		Description:  data.Description,
	}
}

func GetWorkUnitProtoToResponse(data *pb.WorkUnit) *pb.GetWorkUnitResponse {
	return &pb.GetWorkUnitResponse{
		WorkUnit: data,
	}
}

func CreateWorkUnitRequestToModel(data *pb.CreateWorkUnitRequest) *models.WorkUnit {
	return &models.WorkUnit{
		Name:         data.Name,
		DepartmentID: data.DepartmentId,
		Description:  data.Description,
	}
}

func UpdateWorkUnitRequestToModel(data *pb.UpdateWorkUnitRequest) *models.WorkUnit {
	res := models.WorkUnit{
		Name:         data.Name,
		DepartmentID: data.DepartmentId,
		Description:  data.Description,
	}

	res.ID = data.Id
	return &res
}
