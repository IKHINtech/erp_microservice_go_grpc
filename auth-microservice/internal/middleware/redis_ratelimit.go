package middleware

//
// import (
// 	"context"
// 	"fmt"
// 	"time"
//
// 	"github.com/redis/go-redis/v9"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"
// )
//
// type RedisRateLimiter struct {
// 	client *redis.Client
// 	limit  int
// 	window time.Duration
// }
//
// func NewRedisRateLimiter(addr string, limit int, window time.Duration) *RedisRateLimiter {
// 	return &RedisRateLimiter{
// 		client: redis.NewClient(&redis.Options{Addr: addr}),
// 		limit:  limit,
// 		window: window,
// 	}
// }
//
// func (r *RedisRateLimiter) Intercept(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
// 	key := fmt.Sprintf("rate_limit:%s", info.FullMethod)
//
// 	current, err := r.client.Incr(ctx, key).Result()
// 	if err != nil {
// 		return nil, status.Errorf(codes.Internal, "failed to check rate limit")
// 	}
//
// 	if current == 1 {
// 		r.client.Expire(ctx, key, r.window)
// 	}
//
// 	if current > int64(r.limit) {
// 		return nil, status.Errorf(codes.ResourceExhausted, "rate limit exceeded")
// 	}
//
// 	return handler(ctx, req)
// }
