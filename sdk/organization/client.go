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

func (c *OrganizationClient) CreateOrganization(ctx context.Context, org *organizationv1.CreateOrganizationRequest) (*organizationv1.Organization, error) {
	res, err := c.client.CreateOrganization(ctx, org)
	if err != nil {
		return nil, err
	}
	return res, nil
}
