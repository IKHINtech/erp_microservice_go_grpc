package handlers

import (
	"context"
	"errors"

	"github.com/IKHINtech/erp_microservice_go_grpc/auth-microservice/config"
	"github.com/IKHINtech/erp_microservice_go_grpc/auth-microservice/internal/models"
	"github.com/IKHINtech/erp_microservice_go_grpc/auth-microservice/internal/repositories"
	"github.com/IKHINtech/erp_microservice_go_grpc/auth-microservice/internal/utils"
	authv1 "github.com/IKHINtech/erp_microservice_go_grpc/auth-microservice/proto/auth/v1"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type AuthServer struct {
	authv1.UnimplementedAuthServiceServer
	Repo *repositories.AuthRepository
}

func (s *AuthServer) RegisterUser(ctx context.Context, req *authv1.RegisterUserRequest) (*authv1.RegisterUserResponse, error) {
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email dan password harus diisi")
	}
	// 2. Cek apakah user sudah ada
	_, err := s.Repo.GetUserByEmail(req.Email)
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

	err = s.Repo.CreateUser(&user)
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
		return &authv1.ValidateTokenResponse{IsValid: false}, nil
	}

	// Dapatkan roles dari database
	roles, err := s.Repo.GetUserRoles(claims.UserID)
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
	role, err := s.Repo.GetRoleByName(req.RoleName)
	if err != nil {
		return nil, errors.New("role tidak ditemukan")
	}

	// 2. Assign role ke user
	err = s.Repo.AssignRoleToUser(req.UserId, role.ID.String())
	if err != nil {
		return nil, err
	}

	return &authv1.AssignRoleResponse{Success: true}, nil
}

func (s *AuthServer) Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	user, err := s.Repo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, status.Error(codes.NotFound, "user tidak ditemukan")
	}

	// 2. Validasi password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, status.Error(codes.Unauthenticated, "password salah")
	}

	// 3. Generate JWT
	token, err := utils.GenerateJWT(user.ID.String(), config.CFG)
	if err != nil {
		return nil, status.Error(codes.Internal, "gagal generate token")
	}

	return &authv1.LoginResponse{
		Token:  token,
		UserId: user.ID.String(),
	}, nil
}
