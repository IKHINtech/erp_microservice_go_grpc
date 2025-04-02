package organization

import (
	"time"

	"context"
	organizationv1 "github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/pb/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// TODO: lanjut dari sini
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
// GetOrganization(ctx context.Context, in *GetOrganizationRequest, opts ...grpc.CallOption) (*GetOrganizationResponse, error)
// ListOrganization(ctx context.Context, in *ListOrganizationRequest, opts ...grpc.CallOption) (*ListOrganizationResponse, error)
// UpdateOrganization(ctx context.Context, in *UpdateOrganizationRequest, opts ...grpc.CallOption) (*Organization, error)
// DeleteOrganization(ctx context.Context, in *DeleteOrganizationRequest, opts ...grpc.CallOption) (*Organization, error)
// // Department Endpoint
// CreateDepartment(ctx context.Context, in *CreateDepartmentRequest, opts ...grpc.CallOption) (*Department, error)
// GetDepartment(ctx context.Context, in *GetDepartmentRequest, opts ...grpc.CallOption) (*GetDepartmentResponse, error)
// ListDepartment(ctx context.Context, in *ListDepartmentRequest, opts ...grpc.CallOption) (*ListDepartmentResponse, error)
// UpdateDepartment(ctx context.Context, in *UpdateDepartmentRequest, opts ...grpc.CallOption) (*Department, error)
// DeleteDepartment(ctx context.Context, in *DeleteDepartmentRequest, opts ...grpc.CallOption) (*Department, error)
// // WorkUnit Endpoint
// CreateWorkUnit(ctx context.Context, in *CreateWorkUnitRequest, opts ...grpc.CallOption) (*WorkUnit, error)
// GetWorkUnit(ctx context.Context, in *GetWorkUnitRequest, opts ...grpc.CallOption) (*GetWorkUnitResponse, error)
// ListWorkUnit(ctx context.Context, in *ListWorkUnitRequest, opts ...grpc.CallOption) (*ListWorkUnitResponse, error)
// DeleteWorkUnit(ctx context.Context, in *DeleteWorkUnitRequest, opts ...grpc.CallOption) (*WorkUnit, error)
// // Position Endpoint
// CreatePosition(ctx context.Context, in *CreatePositionRequest, opts ...grpc.CallOption) (*Position, error)
// GetPosition(ctx context.Context, in *GetPositionRequest, opts ...grpc.CallOption) (*GetPositionResponse, error)
// ListPosition(ctx context.Context, in *ListPositionRequest, opts ...grpc.CallOption) (*ListPositionResponse, error)
// UpdatePosition(ctx context.Context, in *UpdatePositionRequest, opts ...grpc.CallOption) (*Position, error)
// DeletePosition(ctx context.Context, in *DeletePositionRequest, opts ...grpc.CallOption) (*Position, error)
// // Work Hierarchy Endpoint
// CreatePositionHierarchy(ctx context.Context, in *CreatePositionHierarchyRequest, opts ...grpc.CallOption) (*PositionHierarchy, error)
// GetPositionHierarchyBySuperior(ctx context.Context, in *GetPositionHierarchyBySuperiorRequest, opts ...grpc.CallOption) (*ListPositionHierarchyResponse, error)
// GetPositionHierarchyBySubordinated(ctx context.Context, in *GetPositionHierarchyBySubordinatedRequest, opts ...grpc.CallOption) (*ListPositionHierarchyResponse, error)
// DeletePositionHierarchy(ctx context.Context, in *DeletePositionHierarchyRequest, opts ...grpc.CallOption) (*PositionHierarchy, error)

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
