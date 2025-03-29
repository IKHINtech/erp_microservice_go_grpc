package dto

import (
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/internal/models"
	pb "github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/pb/v1"
)

func GetOrganizationProtoToResponse(data *pb.Organization) *pb.GetOrganizationResponse {
	return &pb.GetOrganizationResponse{
		Organization: data,
	}
}

func OrganizationToProto(organization *models.Organization) *pb.Organization {
	var parrentID string // TODO: nanti ini bakal diubah ke nullable

	if organization.ParentOrganizationID != nil {
		parrentID = *organization.ParentOrganizationID
	}

	return &pb.Organization{
		Id:                   organization.ID,
		Name:                 organization.Name,
		Description:          organization.Description,
		Type:                 organization.Type,
		ParentOrganizationId: parrentID,
	}
}

func CreateOrganizationRequestToModel(organization *pb.CreateOrganizationRequest) *models.Organization {
	return &models.Organization{
		Name:                 organization.Name,
		Description:          organization.Description,
		Type:                 organization.Type,
		ParentOrganizationID: &organization.ParentOrganizationId, // TODO: nanti ini bakal diubah ke nullable
	}
}

func UpdateOrganizationRequestToModel(organization *pb.UpdateOrganizationRequest) *models.Organization {
	data := models.Organization{
		Name:                 organization.Name,
		Description:          organization.Description,
		Type:                 organization.Type,
		ParentOrganizationID: &organization.ParentOrganizationId, // TODO: nanti ini bakal diubah ke nullable
	}
	data.ID = organization.Id
	return &data
}
