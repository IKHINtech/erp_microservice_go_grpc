package repositories

import (
	"github.com/IKHINtech/erp_microservice_go_grpc/auth-microservice/internal/models"
	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
	GetRoleByName(name string) (*models.Role, error)
	AssignRoleToUser(userID, roleID string) error
	GetUsersByRole(roleID string) ([]*models.User, error)
	GetUserRoles(userID string) ([]string, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

// Auto-migrate model ke database
func (r *authRepository) Migrate() error {
	return r.db.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Permission{},
	)
}

// Contoh fungsi: CreateUser
func (r *authRepository) CreateUser(user *models.User) error {
	return r.db.Create(&user).Error
}

func (r *authRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *authRepository) GetUserByID(id string) (*models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (r *authRepository) GetRoleByName(name string) (*models.Role, error) {
	var role models.Role
	err := r.db.Where("name = ?", name).First(&role).Error
	return &role, err
}

func (r *authRepository) AssignRoleToUser(userID, roleID string) error {
	return r.db.Exec("INSERT INTO user_roles (user_id, role_id) VALUES (?, ?)", userID, roleID).Error
}

func (r *authRepository) GetUsersByRole(roleID string) ([]*models.User, error) {
	var users []*models.User
	err := r.db.Joins("JOIN user_roles ON users.id = user_roles.user_id").Where("user_roles.role_id = ?", roleID).Find(&users).Error
	return users, err
}

func (r *authRepository) GetUserRoles(userID string) ([]string, error) {
	var roles []string
	err := r.db.Model(&models.Role{}).
		Joins("JOIN user_roles ON roles.id = user_roles.role_id").
		Where("user_roles.user_id = ?", userID).
		Pluck("roles.name", &roles).
		Error
	return roles, err
}
