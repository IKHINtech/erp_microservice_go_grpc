package middleware

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UnaryServerInterceptor untuk logging
func LoggingUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	start := time.Now()

	// Panggil handler
	resp, err = handler(ctx, req)

	// Log request
	duration := time.Since(start)
	statusCode := codes.Unknown
	if st, ok := status.FromError(err); ok {
		statusCode = st.Code()
	}

	log.Printf("Method: %s | Status: %s | Duration: %v | Request: %+v",
		info.FullMethod,
		statusCode.String(),
		duration,
		req,
	)

	return resp, err
}

// StreamInterceptor (opsional)
func LoggingStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Printf("Stream started: %s", info.FullMethod)
	err := handler(srv, ss)
	log.Printf("Stream ended: %s (error: %v)", info.FullMethod, err)
	return err
}
