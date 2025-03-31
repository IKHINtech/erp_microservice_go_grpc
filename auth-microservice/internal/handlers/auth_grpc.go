package handlers

import (
	"context"
	"errors"
	"time"

	"github.com/IKHINtech/erp_microservice_go_grpc/auth-microservice/config"
	"github.com/IKHINtech/erp_microservice_go_grpc/auth-microservice/internal/models"
	"github.com/IKHINtech/erp_microservice_go_grpc/auth-microservice/internal/repositories"
	"github.com/IKHINtech/erp_microservice_go_grpc/auth-microservice/internal/utils"
	authv1 "github.com/IKHINtech/erp_microservice_go_grpc/auth-microservice/pb/v1"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type AuthServer struct {
	authv1.UnimplementedAuthServiceServer
	repo repositories.AuthRepository
}

func NewAuthServer(repo repositories.AuthRepository) *AuthServer {
	return &AuthServer{
		UnimplementedAuthServiceServer: authv1.UnimplementedAuthServiceServer{},
		repo:                           repo,
	}
}

func (s *AuthServer) RegisterUser(ctx context.Context, req *authv1.RegisterUserRequest) (*authv1.RegisterUserResponse, error) {
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email dan password harus diisi")
	}
	// 2. Cek apakah user sudah ada
	_, err := s.repo.GetUserByEmail(req.Email)
	if err == nil {
		return nil, errors.New("email sudah terdaftar")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// 3. Buat user baru
	user := models.User{
		Email:        req.Email,
		PasswordHash: req.Password, // Akan di-hash otomatis oleh BeforeSave hook
		FullName:     req.FullName,
		IsActive:     true,
	}

	err = s.repo.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	return &authv1.RegisterUserResponse{
		UserId: user.ID.String(),
	}, nil
}

// ValidateToken implementation
func (s *AuthServer) ValidateToken(ctx context.Context, req *authv1.ValidateTokenRequest) (*authv1.ValidateTokenResponse, error) {
	claims, err := utils.ValidateJWT(req.Token, config.CFG.JWTSecret)
	if err != nil {
		return &authv1.ValidateTokenResponse{IsValid: false}, err
	}

	// Dapatkan roles dari database
	roles, err := s.repo.GetUserRoles(claims.UserID)
	if err != nil {
		return nil, err
	}

	return &authv1.ValidateTokenResponse{
		IsValid: true,
		UserId:  claims.UserID,
		Roles:   roles,
	}, nil
}

// AssignRoleToUser implementation
func (s *AuthServer) AssignRoleToUser(ctx context.Context, req *authv1.AssignRoleRequest) (*authv1.AssignRoleResponse, error) {
	// 1. Dapatkan role dari database
	role, err := s.repo.GetRoleByName(req.RoleName)
	if err != nil {
		return nil, errors.New("role tidak ditemukan")
	}

	// 2. Assign role ke user
	err = s.repo.AssignRoleToUser(req.UserId, role.ID.String())
	if err != nil {
		return nil, err
	}

	return &authv1.AssignRoleResponse{Success: true}, nil
}

func (s *AuthServer) Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, status.Error(codes.NotFound, "user tidak ditemukan")
	}

	// 2. Validasi password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, status.Error(codes.Unauthenticated, "password salah")
	}

	accessToken, err := utils.GenerateJWT(user.ID.String(), "access", config.CFG.JWTSecret, 24*time.Hour)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to generate token")
	}

	refreshToken, err := utils.GenerateJWT(user.ID.String(), "refresh", config.CFG.JWTSecret, 7*24*time.Hour)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to generate refresh token")
	}

	return &authv1.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserId:       user.ID.String(),
	}, nil
}

func (s *AuthServer) RefreshToken(ctx context.Context, req *authv1.RefreshTokenRequest) (*authv1.RefreshTokenResponse, error) {
	claims, err := utils.ValidateJWT(req.RefreshToken, config.CFG.JWTSecret)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid refresh token")
	}

	if claims.TokenType != "refresh" {
		return nil, status.Error(codes.Unauthenticated, "not a refresh token")
	}

	newAccessToken, err := utils.GenerateJWT(claims.UserID, "access", config.CFG.JWTSecret, 24*time.Hour)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to generate token")
	}

	return &authv1.RefreshTokenResponse{
		AccessToken:  newAccessToken,
		RefreshToken: req.RefreshToken,
		ExpiresIn:    900,
	}, nil
}
