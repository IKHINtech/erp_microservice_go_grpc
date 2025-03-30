package dto

import (
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/internal/models"
	pb "github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/pb/v1"
)

func PositionToProto(data *models.Position) *pb.Position {
	return &pb.Position{
		Id:          data.ID,
		Name:        data.Name,
		Description: data.Description,
	}
}

func GetPositionProtoToResponse(data *pb.Position) *pb.GetPositionResponse {
	return &pb.GetPositionResponse{
		Position: data,
	}
}

func CreatePositionRequestToModel(data *pb.CreatePositionRequest) *models.Position {
	return &models.Position{
		Name:        data.Name,
		Description: data.Description,
	}
}

func UpdatePositionRequestToModel(data *pb.UpdatePositionRequest) *models.Position {
	res := models.Position{
		Name:        data.Name,
		Description: data.Description,
	}

	res.ID = data.Id
	return &res
}
