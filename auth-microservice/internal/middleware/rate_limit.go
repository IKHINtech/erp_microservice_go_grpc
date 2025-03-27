package middleware

import (
	"context"
	"time"

	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RateLimitInterceptor(limit int, interval time.Duration) grpc.UnaryServerInterceptor {
	limiter := rate.NewLimiter(rate.Every(interval), limit)

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if !limiter.Allow() {
			return nil, status.Errorf(codes.ResourceExhausted, "too many requests")
		}
		return handler(ctx, req)
	}
}
