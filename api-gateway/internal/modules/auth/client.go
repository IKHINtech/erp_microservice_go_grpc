package auth

import (
	"context"

	authv1 "github.com/IKHINtech/erp_microservice_go_grpc/gen-proto/proto/auth/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthClient struct {
	client authv1.AuthServiceClient
}

func NewAuthClient(addr string) (*AuthClient, error) {
	conn, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	return &AuthClient{
		authv1.NewAuthServiceClient(conn),
	}, nil
}

func (c *AuthClient) Login(ctx context.Context, email, password string) (*authv1.LoginResponse, error) {
	res, err := c.client.Login(ctx, &authv1.LoginRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *AuthClient) Register(ctx context.Context, req *authv1.RegisterUserRequest) (*authv1.RegisterUserResponse, error) {
	res, err := c.client.RegisterUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (c *AuthClient) ValidateToken(ctx context.Context, token string) (*authv1.ValidateTokenResponse, error) {
	res, err := c.client.ValidateToken(ctx, &authv1.ValidateTokenRequest{
		Token: token,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AuthClient) AssignRoleToUser(ctx context.Context, req *authv1.AssignRoleRequest) (*authv1.AssignRoleResponse, error) {
	res, err := c.client.AssignRoleToUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
