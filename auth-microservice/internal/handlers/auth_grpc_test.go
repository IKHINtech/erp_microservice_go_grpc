package handlers_test

import (
	"context"
	"testing"

	"github.com/IKHINtech/erp_microservice_go_grpc/auth-microservice/internal/handlers"
	"github.com/IKHINtech/erp_microservice_go_grpc/auth-microservice/internal/models"
	authv1 "github.com/IKHINtech/erp_microservice_go_grpc/auth-microservice/proto/auth/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type MockAuthRepository struct {
	mock.Mock
}

func (m *MockAuthRepository) AssignRoleToUser(userID, roleID string) error {
	args := m.Called(userID, roleID)
	return args.Error(0)
}

func (m *MockAuthRepository) Migrate() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockAuthRepository) GetUserByEmail(email string) (*models.User, error) {
	args := m.Called(email)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockAuthRepository) GetUserByID(id string) (*models.User, error) {
	args := m.Called(id)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockAuthRepository) GetRoleByName(name string) (*models.Role, error) {
	args := m.Called(name)
	return args.Get(0).(*models.Role), args.Error(1)
}

func (m *MockAuthRepository) GetUsersByRole(roleID string) ([]*models.User, error) {
	args := m.Called(roleID)
	return args.Get(0).([]*models.User), args.Error(1)
}

func (m *MockAuthRepository) GetUserRoles(userID string) ([]string, error) {
	args := m.Called(userID)
	return args.Get(0).([]string), args.Error(1)
}

func (m *MockAuthRepository) CreateUser(data *models.User) error {
	args := m.Called(data.Email, data.PasswordHash)
	return args.Error(1)
}

func TestRegisterUser(t *testing.T) {
	mockRepo := new(MockAuthRepository)
	handler := handlers.NewAuthServer(mockRepo)
	t.Run("Success", func(t *testing.T) {
		mockRepo.On("CreateUser", "email@example.com", "password").Return(&models.User{}, nil)
		res, err := handler.RegisterUser(nil, &authv1.RegisterUserRequest{Email: "email@example1.com", Password: "password"})
		assert.NoError(t, err)
		assert.Equal(t, "email", res.UserId)
		mockRepo.AssertExpectations(t)
	})
	t.Run("Invalid Email", func(t *testing.T) {
		_, err := handler.RegisterUser(context.Background(), &authv1.RegisterUserRequest{
			Email:    "",
			Password: "password123",
		})

		assert.Equal(t, codes.InvalidArgument, status.Code(err))
	})
}

func TestRegisterUser1(t *testing.T) {
	mockRepo := new(MockAuthRepository)
	handler := handlers.NewAuthServer(mockRepo)

	t.Run("Success", func(t *testing.T) {
		// Buat mock untuk GetUserByEmail
		mockRepo.On("GetUserByEmail", "email@example1.com").Return(nil, gorm.ErrRecordNotFound)

		// Buat mock untuk CreateUser
		user := &models.User{
			Email:        "email@example1.com",
			PasswordHash: "password",
		}
		mockRepo.On("CreateUser", user).Return(nil)

		// Panggil handler RegisterUser dengan request yang sesuai
		res, err := handler.RegisterUser(nil, &authv1.RegisterUserRequest{Email: "email@example1.com", Password: "password"})

		// Pastikan tidak ada error dan userId sesuai
		assert.NoError(t, err)
		assert.Equal(t, "email@example1.com", res.UserId)
	})

	t.Run("UserAlreadyExists", func(t *testing.T) {
		// Kasus di mana email sudah terdaftar
		mockRepo.On("GetUserByEmail", "email@example1.com").Return(&models.User{Email: "email@example1.com"}, nil)

		// Panggil handler RegisterUser dengan request yang sesuai
		_, err := handler.RegisterUser(nil, &authv1.RegisterUserRequest{Email: "email@example1.com", Password: "password"})

		// Pastikan error sesuai dengan skenario "user already exists"
		assert.Error(t, err)
	})
}

func TestLogin(t *testing.T) {
	mockRepo := new(MockAuthRepository)
	handler := handlers.NewAuthServer(mockRepo)

	tests := []struct {
		name    string
		req     *authv1.LoginRequest
		wantErr bool
		code    codes.Code
	}{
		{
			name:    "Empty Password",
			req:     &authv1.LoginRequest{Email: "test@example.com", Password: ""},
			wantErr: true,
			code:    codes.InvalidArgument,
		},
		{
			name:    "Valid Request",
			req:     &authv1.LoginRequest{Email: "valid@example.com", Password: "password123"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := handler.Login(context.Background(), tt.req)
			if tt.wantErr {
				assert.Equal(t, tt.code, status.Code(err))
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func BenchmarkRegisterUser(b *testing.B) {
	mockRepo := new(MockAuthRepository)
	handler := handlers.NewAuthServer(mockRepo)

	req := &authv1.RegisterUserRequest{
		Email:    "bench@example.com",
		Password: "password123",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = handler.RegisterUser(context.Background(), req)
	}
}
