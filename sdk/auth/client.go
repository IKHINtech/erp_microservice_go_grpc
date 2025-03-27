package auth

import (
	"context"
	"time"

	authv1 "github.com/IKHINtech/erp_microservice_go_grpc/auth-microservice/proto/auth/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn    *grpc.ClientConn
	timeout time.Duration
}

func NewClient(addr string, timeoutSec int) (*Client, error) {
	conn, err := grpc.Dial(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(10*1024*1024)),
	)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn:    conn,
		timeout: time.Duration(timeoutSec) * time.Second,
	}, nil
}

// Contoh method untuk Login
func (c *Client) Login(ctx context.Context, email, password string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	client := authv1.NewAuthServiceClient(c.conn)
	res, err := client.Login(ctx, &authv1.LoginRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return "", err
	}

	return res.Token, nil
}
