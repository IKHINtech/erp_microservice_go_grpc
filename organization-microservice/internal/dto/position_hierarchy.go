package dto

import (
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/internal/models"
	pb "github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/pb/v1"
)

func PositionHierarchyToProto(data *models.PositionHierarchy) *pb.PositionHierarchy {
	return &pb.PositionHierarchy{
		SuperiorPositionId:     data.SuperiorPositionID,
		SubordinatedPositionId: data.SubordinatedPositionID,
	}
}

func GetPositionHierarchyProtoToResponse(data []*pb.PositionHierarchy) *pb.GetPositionHierarchyResponse {
	return &pb.GetPositionHierarchyResponse{
		PositionHierarchy: data,
	}
}

func CreatePositionHierarchyRequestToModel(data *pb.CreatePositionHierarchyRequest) *models.PositionHierarchy {
	return &models.PositionHierarchy{
		SuperiorPositionID:     data.SuperiorPositionId,
		SubordinatedPositionID: data.SubordinatedPositionId,
	}
}
