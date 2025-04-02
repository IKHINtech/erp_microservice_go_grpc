package server

import (
	"log"
	"net"
	"time"

	"github.com/IKHINtech/erp_microservice_go_grpc/auth-microservice/config"
	"github.com/IKHINtech/erp_microservice_go_grpc/auth-microservice/internal/handlers"
	"github.com/IKHINtech/erp_microservice_go_grpc/auth-microservice/internal/middleware"
	"github.com/IKHINtech/erp_microservice_go_grpc/auth-microservice/internal/repositories"
	authv1 "github.com/IKHINtech/erp_microservice_go_grpc/auth-microservice/pb/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

func StartServer(db *gorm.DB) {

	// Local rate limiting (10 requests/second)
	rateLimit := middleware.RateLimitInterceptor(10, time.Second)

	authRepo := repositories.NewAuthRepository(db)
	if err := authRepo.Migrate(); err != nil {
		log.Fatal("Gagal migrate database:", err)
	}

	lis, err := net.Listen("tcp", "0.0.0.0:"+config.CFG.AuthPort) // Ganti localhost ke 0.0.0.0
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	authServer := handlers.NewAuthServer(authRepo)

	s := grpc.NewServer(
		grpc.StreamInterceptor(middleware.LoggingStreamInterceptor),
		grpc.ChainUnaryInterceptor(
			middleware.LoggingUnaryInterceptor,
			rateLimit,
		),
	)
	authv1.RegisterAuthServiceServer(s, authServer)

	reflection.Register(s)

	log.Printf("Server running on %s", lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
