package repositories

import (
	"github.com/IKHINtech/erp_microservice_go_grpc/auth-microservice/internal/models"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

// Auto-migrate model ke database
func (r *AuthRepository) Migrate() error {
	return r.db.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Permission{},
	)
}

// Contoh fungsi: CreateUser
func (r *AuthRepository) CreateUser(email, password string) (*models.User, error) {
	user := models.User{
		Email:        email,
		PasswordHash: password, // Akan di-hash oleh BeforeSave hook
	}
	err := r.db.Create(&user).Error
	return &user, err
}
