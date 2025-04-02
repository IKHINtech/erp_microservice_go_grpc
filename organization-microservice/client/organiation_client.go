package organization

import (
	"time"

	"context"
	organizationv1 "github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/pb/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OrganizationClient struct {
	conn    *grpc.ClientConn
	timeOut time.Duration
}

func NewOrganizationClient(addr string, timeoutSec int) (*OrganizationClient, error) {
	conn, err := grpc.NewClient(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(10*1024*1024)),
	)
	if err != nil {
		return nil, err
	}
	return &OrganizationClient{
		conn:    conn,
		timeOut: time.Duration(timeoutSec) * time.Second,
	}, nil
}

// CreateOrganization(ctx context.Context, in *CreateOrganizationRequest, opts ...grpc.CallOption) (*Organization, error)
func (c *OrganizationClient) CreateOrganization(ctx context.Context, org *organizationv1.CreateOrganizationRequest) (*organizationv1.Organization, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeOut)
	defer cancel()
	client := organizationv1.NewOrganizationServiceClient(c.conn)
	res, err := client.CreateOrganization(ctx, org)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetOrganization(ctx context.Context, in *GetOrganizationRequest, opts ...grpc.CallOption) (*GetOrganizationResponse, error)
func (c *OrganizationClient) GetOrganization(ctx context.Context, req *organizationv1.GetOrganizationRequest) (*organizationv1.GetOrganizationResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeOut)
	defer cancel()
	client := organizationv1.NewOrganizationServiceClient(c.conn)
	res, err := client.GetOrganization(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, err
}

// ListOrganization(ctx context.Context, in *ListOrganizationRequest, opts ...grpc.CallOption) (*ListOrganizationResponse, error)
func (c *OrganizationClient) ListOrganization(ctx context.Context) (*organizationv1.ListOrganizationResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeOut)
	defer cancel()
	client := organizationv1.NewOrganizationServiceClient(c.conn)
	res, err := client.ListOrganization(ctx, &organizationv1.ListOrganizationRequest{})
	return res, err
}

// UpdateOrganization(ctx context.Context, in *UpdateOrganizationRequest, opts ...grpc.CallOption) (*Organization, error)
func (c *OrganizationClient) UpdateOrganization(ctx context.Context, org *organizationv1.UpdateOrganizationRequest) (*organizationv1.Organization, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeOut)
	defer cancel()
	client := organizationv1.NewOrganizationServiceClient(c.conn)
	res, err := client.UpdateOrganization(ctx, org)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DeleteOrganization(ctx context.Context, in *DeleteOrganizationRequest, opts ...grpc.CallOption) (*Organization, error)
func (c *OrganizationClient) DeleteOrganization(ctx context.Context, org *organizationv1.DeleteOrganizationRequest) (*organizationv1.Organization, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeOut)
	defer cancel()
	client := organizationv1.NewOrganizationServiceClient(c.conn)
	res, err := client.DeleteOrganization(ctx, org)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// // Department Endpoint
// CreateDepartment(ctx context.Context, in *CreateDepartmentRequest, opts ...grpc.CallOption) (*Department, error)
func (c *OrganizationClient) CreateDepartment(ctx context.Context, org *organizationv1.CreateDepartmentRequest) (*organizationv1.Department, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeOut)
	defer cancel()
	client := organizationv1.NewOrganizationServiceClient(c.conn)
	res, err := client.CreateDepartment(ctx, org)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetDepartment(ctx context.Context, in *GetDepartmentRequest, opts ...grpc.CallOption) (*GetDepartmentResponse, error)
func (c *OrganizationClient) GetDepartment(ctx context.Context, org *organizationv1.GetDepartmentRequest) (*organizationv1.GetDepartmentResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeOut)
	defer cancel()
	client := organizationv1.NewOrganizationServiceClient(c.conn)
	res, err := client.GetDepartment(ctx, org)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ListDepartment(ctx context.Context, in *ListDepartmentRequest, opts ...grpc.CallOption) (*ListDepartmentResponse, error)
func (c *OrganizationClient) ListDepartment(ctx context.Context) (*organizationv1.ListDepartmentResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeOut)
	defer cancel()
	client := organizationv1.NewOrganizationServiceClient(c.conn)
	res, err := client.ListDepartment(ctx, &organizationv1.ListDepartmentRequest{})
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UpdateDepartment(ctx context.Context, in *UpdateDepartmentRequest, opts ...grpc.CallOption) (*Department, error)
func (c *OrganizationClient) UpdateDepartment(ctx context.Context, org *organizationv1.UpdateDepartmentRequest) (*organizationv1.Department, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeOut)
	defer cancel()
	client := organizationv1.NewOrganizationServiceClient(c.conn)
	res, err := client.UpdateDepartment(ctx, org)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DeleteDepartment(ctx context.Context, in *DeleteDepartmentRequest, opts ...grpc.CallOption) (*Department, error)
func (c *OrganizationClient) DeleteDepartment(ctx context.Context, org *organizationv1.DeleteDepartmentRequest) (*organizationv1.Department, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeOut)
	defer cancel()
	client := organizationv1.NewOrganizationServiceClient(c.conn)
	res, err := client.DeleteDepartment(ctx, org)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// // WorkUnit Endpoint
// CreateWorkUnit(ctx context.Context, in *CreateWorkUnitRequest, opts ...grpc.CallOption) (*WorkUnit, error)
func (c *OrganizationClient) CreateWorkUnit(ctx context.Context, org *organizationv1.CreateWorkUnitRequest) (*organizationv1.WorkUnit, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeOut)
	defer cancel()
	client := organizationv1.NewOrganizationServiceClient(c.conn)
	res, err := client.CreateWorkUnit(ctx, org)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetWorkUnit(ctx context.Context, in *GetWorkUnitRequest, opts ...grpc.CallOption) (*GetWorkUnitResponse, error)
func (c *OrganizationClient) GetWorkUnit(ctx context.Context, org *organizationv1.GetWorkUnitRequest) (*organizationv1.GetWorkUnitResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeOut)
	defer cancel()
	client := organizationv1.NewOrganizationServiceClient(c.conn)
	res, err := client.GetWorkUnit(ctx, org)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ListWorkUnit(ctx context.Context, in *ListWorkUnitRequest, opts ...grpc.CallOption) (*ListWorkUnitResponse, error)
func (c *OrganizationClient) ListWorkUnit(ctx context.Context) (*organizationv1.ListWorkUnitResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeOut)
	defer cancel()
	client := organizationv1.NewOrganizationServiceClient(c.conn)
	res, err := client.ListWorkUnit(ctx, &organizationv1.ListWorkUnitRequest{})
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DeleteWorkUnit(ctx context.Context, in *DeleteWorkUnitRequest, opts ...grpc.CallOption) (*WorkUnit, error)
func (c *OrganizationClient) DeleteWorkUnit(ctx context.Context, org *organizationv1.DeleteWorkUnitRequest) (*organizationv1.WorkUnit, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeOut)
	defer cancel()
	client := organizationv1.NewOrganizationServiceClient(c.conn)
	res, err := client.DeleteWorkUnit(ctx, org)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// // Position Endpoint
// CreatePosition(ctx context.Context, in *CreatePositionRequest, opts ...grpc.CallOption) (*Position, error)
func (c *OrganizationClient) CreatePosition(ctx context.Context, org *organizationv1.CreatePositionRequest) (*organizationv1.Position, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeOut)
	defer cancel()
	client := organizationv1.NewOrganizationServiceClient(c.conn)
	res, err := client.CreatePosition(ctx, org)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetPosition(ctx context.Context, in *GetPositionRequest, opts ...grpc.CallOption) (*GetPositionResponse, error)
func (c *OrganizationClient) GetPosition(ctx context.Context, org *organizationv1.GetPositionRequest) (*organizationv1.GetPositionResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeOut)
	defer cancel()
	client := organizationv1.NewOrganizationServiceClient(c.conn)
	res, err := client.GetPosition(ctx, org)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//TODO: lanjut disini

// ListPosition(ctx context.Context, in *ListPositionRequest, opts ...grpc.CallOption) (*ListPositionResponse, error)
func (c *OrganizationClient) ListPosition(ctx context.Context) (*organizationv1.ListOrganizationResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeOut)
	defer cancel()
	client := organizationv1.NewOrganizationServiceClient(c.conn)
	res, err := client.ListOrganization(ctx, &organizationv1.ListOrganizationRequest{})
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UpdatePosition(ctx context.Context, in *UpdatePositionRequest, opts ...grpc.CallOption) (*Position, error)
func (c *OrganizationClient) UpdatePosition(ctx context.Context, org *organizationv1.UpdateOrganizationRequest) (*organizationv1.Organization, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeOut)
	defer cancel()
	client := organizationv1.NewOrganizationServiceClient(c.conn)
	res, err := client.UpdateOrganization(ctx, org)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DeletePosition(ctx context.Context, in *DeletePositionRequest, opts ...grpc.CallOption) (*Position, error)
func (c *OrganizationClient) DeletePosition(ctx context.Context, org *organizationv1.DeleteOrganizationRequest) (*organizationv1.Organization, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeOut)
	defer cancel()
	client := organizationv1.NewOrganizationServiceClient(c.conn)
	res, err := client.DeleteOrganization(ctx, org)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// // Work Hierarchy Endpoint
// CreatePositionHierarchy(ctx context.Context, in *CreatePositionHierarchyRequest, opts ...grpc.CallOption) (*PositionHierarchy, error)
func (c *OrganizationClient) CreatePositionHierarchy(ctx context.Context, org *organizationv1.CreateOrganizationRequest) (*organizationv1.Organization, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeOut)
	defer cancel()
	client := organizationv1.NewOrganizationServiceClient(c.conn)
	res, err := client.CreateOrganization(ctx, org)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetPositionHierarchyBySuperior(ctx context.Context, in *GetPositionHierarchyBySuperiorRequest, opts ...grpc.CallOption) (*ListPositionHierarchyResponse, error)
func (c *OrganizationClient) GetPositionHierarchyBySuperior(ctx context.Context) (*organizationv1.ListOrganizationResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeOut)
	defer cancel()
	client := organizationv1.NewOrganizationServiceClient(c.conn)
	res, err := client.ListOrganization(ctx, &organizationv1.ListOrganizationRequest{})
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetPositionHierarchyBySubordinated(ctx context.Context, in *GetPositionHierarchyBySubordinatedRequest, opts ...grpc.CallOption) (*ListPositionHierarchyResponse, error)
func (c *OrganizationClient) GetPositionHierarchyBySubordinated(ctx context.Context) (*organizationv1.ListOrganizationResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeOut)
	defer cancel()
	client := organizationv1.NewOrganizationServiceClient(c.conn)
	res, err := client.ListOrganization(ctx, &organizationv1.ListOrganizationRequest{})
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DeletePositionHierarchy(ctx context.Context, in *DeletePositionHierarchyRequest, opts ...grpc.CallOption) (*PositionHierarchy, error)
func (c *OrganizationClient) DeletePositionHierarchy(ctx context.Context, org *organizationv1.DeleteOrganizationRequest) (*organizationv1.Organization, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeOut)
	defer cancel()
	client := organizationv1.NewOrganizationServiceClient(c.conn)
	res, err := client.DeleteOrganization(ctx, org)
	if err != nil {
		return nil, err
	}
	return res, nil
}
