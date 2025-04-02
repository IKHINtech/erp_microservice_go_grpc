package auth

import (
	"context"
	"time"

	authv1 "github.com/IKHINtech/erp_microservice_go_grpc/auth-microservice/pb/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthClient struct {
	conn    *grpc.ClientConn
	timeout time.Duration
}

func NewClient(addr string, timeoutSec int) (*AuthClient, error) {
	conn, err := grpc.NewClient(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(10*1024*1024)),
	)
	if err != nil {
		return nil, err
	}

	return &AuthClient{
		conn:    conn,
		timeout: time.Duration(timeoutSec) * time.Second,
	}, nil
}

func (c *AuthClient) Login(ctx context.Context, email, password string) (*authv1.LoginResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	client := authv1.NewAuthServiceClient(c.conn)
	res, err := client.Login(ctx, &authv1.LoginRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *AuthClient) Register(ctx context.Context, req *authv1.RegisterUserRequest) (*authv1.RegisterUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	client := authv1.NewAuthServiceClient(c.conn)
	res, err := client.RegisterUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (c *AuthClient) ValidateToken(ctx context.Context, token string) (*authv1.ValidateTokenResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	client := authv1.NewAuthServiceClient(c.conn)
	res, err := client.ValidateToken(ctx, &authv1.ValidateTokenRequest{
		Token: token,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AuthClient) AssignRoleToUser(ctx context.Context, req *authv1.AssignRoleRequest) (*authv1.AssignRoleResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	client := authv1.NewAuthServiceClient(c.conn)
	res, err := client.AssignRoleToUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
