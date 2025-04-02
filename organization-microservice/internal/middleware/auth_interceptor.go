package middleware

import (
	"context"
	"fmt"
	"strings"

	authSdk "github.com/IKHINtech/erp_microservice_go_grpc/auth-microservice/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// AuthInterceptor adalah gRPC middleware untuk validasi token
func AuthInterceptor(authClient *authSdk.AuthClient) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		// Ambil token dari metadata (header gRPC)
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(401, "metadata not found")
		}

		authorization := md["authorization"]
		if len(authorization) == 0 {
			return nil, status.Errorf(401, "token not provided")
		}

		// Ambil nilai pertama dari slice
		authHeader := authorization[0]

		// Pisahkan "Bearer" dari token
		parts := strings.SplitN(authHeader, " ", 2)

		// Validasi apakah formatnya benar
		if len(parts) != 2 || parts[0] != "Bearer" {
			fmt.Println("Invalid Authorization header format")

			return nil, status.Errorf(401, "invalid authorization header format")
		}

		// Ambil token
		token := parts[1]
		fmt.Println("Extracted Token:", token)

		// Panggil auth-microservice untuk validasi token
		res, err := authClient.ValidateToken(ctx, token)
		if err != nil || !res.IsValid {
			return nil, status.Errorf(401, "invalid token")
		}

		// Lanjutkan request jika token valid
		return handler(ctx, req)
	}
}
