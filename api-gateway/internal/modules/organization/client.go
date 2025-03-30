package organization

import (
	"context"

	organizationv1 "github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/pb/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OrganizationClient struct {
	client organizationv1.OrganizationServiceClient
}

func NewOrganizationClient(addr string) (*OrganizationClient, error) {
	conn, err := grpc.NewClient(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	return &OrganizationClient{
		organizationv1.NewOrganizationServiceClient(conn),
	}, nil
}

func (c *OrganizationClient) GetOrganization(ctx context.Context, id string) (*organizationv1.GetOrganizationResponse, error) {
	res, err := c.client.GetOrganization(ctx, &organizationv1.GetOrganizationRequest{
		Id: id,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *OrganizationClient) GetOrganizations(ctx context.Context) (*organizationv1.ListOrganizationResponse, error) {
	res, err := c.client.ListOrganization(ctx, &organizationv1.ListOrganizationRequest{})
	if err != nil {
		return nil, err
	}
	return res, nil
}
