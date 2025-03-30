package repository

import (
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/internal/models"
	"gorm.io/gorm"
)

type PositionRepository interface {
	CreatePosition(position *models.Position) (*models.Position, error)
	GetPosition(id string) (*models.Position, error)
	GetListPosition() ([]*models.Position, error)
	UpdatePosition(id string, position *models.Position) (*models.Position, error)
	DeletePosition(id string) error
}

type positionRepositoryImpl struct {
	db *gorm.DB
}

func NewPositionRepository(db *gorm.DB) PositionRepository {
	return &positionRepositoryImpl{
		db: db,
	}
}

func (r *positionRepositoryImpl) GetListPosition() ([]*models.Position, error) {
	var position []*models.Position
	err := r.db.Find(&position).Error
	return position, err
}

func (r *positionRepositoryImpl) CreatePosition(position *models.Position) (*models.Position, error) {
	if err := r.db.Create(position).Error; err != nil {
		return nil, err
	}
	return position, nil
}

func (r *positionRepositoryImpl) GetPosition(id string) (*models.Position, error) {
	var position models.Position
	err := r.db.Where("id = ?", id).First(&position).Error
	return &position, err
}

func (r *positionRepositoryImpl) UpdatePosition(id string, position *models.Position) (*models.Position, error) {
	err := r.db.Model(&models.Position{}).Where("id = ?", id).Updates(position).Error
	return position, err
}

func (r *positionRepositoryImpl) DeletePosition(id string) error {
	return r.db.Delete(&models.Position{}, "id = ?", id).Error
}
