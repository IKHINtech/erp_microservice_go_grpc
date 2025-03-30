package repository

import (
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/internal/models"
	"gorm.io/gorm"
)

type DepartmentRepository interface {
	CreateDepartment(department *models.Department) (*models.Department, error)
	GetDepartment(id string) (*models.Department, error)
	GetListDepartment() ([]*models.Department, error)
	UpdateDepartment(id string, department *models.Department) (*models.Department, error)
	DeleteDepartment(id string) error
}

type departmentRepositoryImpl struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) DepartmentRepository {
	return &departmentRepositoryImpl{
		db: db,
	}
}

func (r *departmentRepositoryImpl) GetListDepartment() ([]*models.Department, error) {
	var department []*models.Department
	err := r.db.Find(&department).Error
	return department, err
}

func (r *departmentRepositoryImpl) CreateDepartment(department *models.Department) (*models.Department, error) {
	if err := r.db.Create(department).Error; err != nil {
		return nil, err
	}
	return department, nil
}

func (r *departmentRepositoryImpl) GetDepartment(id string) (*models.Department, error) {
	var department models.Department
	err := r.db.Where("id = ?", id).First(&department).Error
	return &department, err
}

func (r *departmentRepositoryImpl) UpdateDepartment(id string, department *models.Department) (*models.Department, error) {
	err := r.db.Model(&models.Department{}).Where("id = ?", id).Updates(department).Error
	return department, err
}

func (r *departmentRepositoryImpl) DeleteDepartment(id string) error {
	return r.db.Delete(&models.Department{}, "id = ?", id).Error
}
