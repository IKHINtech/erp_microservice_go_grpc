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

	return &pb.Organization{
		Id:                   organization.ID,
		Name:                 organization.Name,
		Description:          organization.Description,
		Type:                 organization.Type,
		ParentOrganizationId: organization.ParentOrganizationID,
	}
}

func CreateOrganizationRequestToModel(organization *pb.CreateOrganizationRequest) *models.Organization {
	return &models.Organization{
		Name:                 organization.Name,
		Description:          organization.Description,
		Type:                 organization.Type,
		ParentOrganizationID: organization.ParentOrganizationId,
	}
}

func UpdateOrganizationRequestToModel(organization *pb.UpdateOrganizationRequest) *models.Organization {
	data := models.Organization{
		Name:                 organization.Name,
		Description:          organization.Description,
		Type:                 organization.Type,
		ParentOrganizationID: organization.ParentOrganizationId,
	}
	data.ID = organization.Id
	return &data
}
