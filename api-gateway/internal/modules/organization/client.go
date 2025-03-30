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

func (c *OrganizationClient) CreateOrganization(ctx context.Context, org *organizationv1.CreateOrganizationRequest) (*organizationv1.Organization, error) {
	res, err := c.client.CreateOrganization(ctx, org)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *OrganizationClient) UpdateOrganization(ctx context.Context, org *organizationv1.UpdateOrganizationRequest) (*organizationv1.Organization, error) {
	res, err := c.client.UpdateOrganization(ctx, org)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *OrganizationClient) DeleteOrganization(ctx context.Context, org *organizationv1.DeleteOrganizationRequest) (*organizationv1.Organization, error) {
	res, err := c.client.DeleteOrganization(ctx, org)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *OrganizationClient) CreateDepartment(ctx context.Context, data CreateDepartmentRequest) (*organizationv1.Department, error) {
	dep := organizationv1.CreateDepartmentRequest{
		Name:        data.Name,
		Description: data.Description,
	}
	res, err := c.client.CreateDepartment(ctx, &dep)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *OrganizationClient) UpdateDepartment(ctx context.Context, dep *organizationv1.UpdateDepartmentRequest) (*organizationv1.Department, error) {
	res, err := c.client.UpdateDepartment(ctx, dep)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *OrganizationClient) DeleteDepartment(ctx context.Context, dep *organizationv1.DeleteDepartmentRequest) (*organizationv1.Department, error) {
	res, err := c.client.DeleteDepartment(ctx, dep)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *OrganizationClient) GetDepartments(ctx context.Context) (*organizationv1.ListDepartmentResponse, error) {
	res, err := c.client.ListDepartment(ctx, &organizationv1.ListDepartmentRequest{})
	if err != nil {
		return nil, err
	}
	return res, nil
}
