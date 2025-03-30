package repository

import (
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/internal/models"
	"gorm.io/gorm"
)

type PositionHierarchyRepository interface {
	CreatePositionHierarchy(positionHierarchy *models.PositionHierarchy) (*models.PositionHierarchy, error)
	GetPositionHierarchy(id string) (*models.PositionHierarchy, error)
	GetListPositionHierarchy(superiorID *string, subordinatedID *string) ([]*models.PositionHierarchy, error)
	UpdatePositionHierarchy(id string, positionHierarchy *models.PositionHierarchy) (*models.PositionHierarchy, error)
	DeletePositionHierarchy(id string) error
}

type positionHierarchyRepositoryImpl struct {
	db *gorm.DB
}

func NewPositionHierarchyRepository(db *gorm.DB) PositionHierarchyRepository {
	return &positionHierarchyRepositoryImpl{
		db: db,
	}
}

func (r *positionHierarchyRepositoryImpl) GetListPositionHierarchy(superiorID *string, subordinatedID *string) ([]*models.PositionHierarchy, error) {
	var positionHierarchy []*models.PositionHierarchy
	query := r.db
	if superiorID != nil {
		query = query.Where("superior_position_id = ?", superiorID)
	}
	if subordinatedID != nil {
		query = query.Where("subordinated_position_id = ?", subordinatedID)
	}
	err := query.Find(&positionHierarchy).Error
	return positionHierarchy, err
}

func (r *positionHierarchyRepositoryImpl) CreatePositionHierarchy(positionHierarchy *models.PositionHierarchy) (*models.PositionHierarchy, error) {
	if err := r.db.Create(positionHierarchy).Error; err != nil {
		return nil, err
	}
	return positionHierarchy, nil
}

func (r *positionHierarchyRepositoryImpl) GetPositionHierarchy(id string) (*models.PositionHierarchy, error) {
	var positionHierarchy models.PositionHierarchy
	err := r.db.Where("id = ?", id).First(&positionHierarchy).Error
	return &positionHierarchy, err
}

func (r *positionHierarchyRepositoryImpl) UpdatePositionHierarchy(id string, positionHierarchy *models.PositionHierarchy) (*models.PositionHierarchy, error) {
	err := r.db.Model(&models.PositionHierarchy{}).Where("id = ?", id).Updates(positionHierarchy).Error
	return positionHierarchy, err
}

func (r *positionHierarchyRepositoryImpl) DeletePositionHierarchy(id string) error {
	return r.db.Delete(&models.PositionHierarchy{}, "id = ?", id).Error
}
