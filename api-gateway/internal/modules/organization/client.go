package client

import (
	organizationv1 "github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/pb/v1"
)

type OrganizationClient struct {
	client organizationv1.OrganizationServiceClient
}
