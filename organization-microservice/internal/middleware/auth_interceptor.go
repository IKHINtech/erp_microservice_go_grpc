package middleware

import (
	"context"
	"log"

	authv1 "github.com/IKHINtech/erp_microservice_go_grpc/gen-proto/proto/auth/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// AuthInterceptor adalah gRPC middleware untuk validasi token
func AuthInterceptor(authClient authv1.AuthServiceClient) grpc.UnaryServerInterceptor {
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

		token := md["authorization"]
		if len(token) == 0 {
			return nil, status.Errorf(401, "token not provided")
		}

		// Panggil auth-microservice untuk validasi token
		res, err := authClient.ValidateToken(ctx, &authv1.ValidateTokenRequest{Token: token[0]})
		log.Println("res", res)
		log.Println("err", err)
		//TODO: pastikan lagi kenapa tidak ada return nya
		if err != nil || !res.IsValid {
			return nil, status.Errorf(401, "invalid token")
		}

		// Lanjutkan request jika token valid
		return handler(ctx, req)
	}
}
