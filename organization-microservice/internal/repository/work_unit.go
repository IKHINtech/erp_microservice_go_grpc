package repository

import (
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/internal/models"
	"gorm.io/gorm"
)

type WorkUnitRepository interface {
	CreateWorkUnit(workUnit *models.WorkUnit) (*models.WorkUnit, error)
	GetWorkUnit(id string) (*models.WorkUnit, error)
	GetListWorkUnit() ([]*models.WorkUnit, error)
	UpdateWorkUnit(id string, workUnit *models.WorkUnit) (*models.WorkUnit, error)
	DeleteWorkUnit(id string) error
}

type workUnitRepositoryImpl struct {
	db *gorm.DB
}

func NewWorkUnitRepository(db *gorm.DB) WorkUnitRepository {
	return &workUnitRepositoryImpl{
		db: db,
	}
}

func (r *workUnitRepositoryImpl) GetListWorkUnit() ([]*models.WorkUnit, error) {
	var workUnit []*models.WorkUnit
	err := r.db.Find(&workUnit).Error
	return workUnit, err
}

func (r *workUnitRepositoryImpl) CreateWorkUnit(workUnit *models.WorkUnit) (*models.WorkUnit, error) {
	if err := r.db.Create(workUnit).Error; err != nil {
		return nil, err
	}
	return workUnit, nil
}

func (r *workUnitRepositoryImpl) GetWorkUnit(id string) (*models.WorkUnit, error) {
	var workUnit models.WorkUnit
	err := r.db.Where("id = ?", id).First(&workUnit).Error
	return &workUnit, err
}

func (r *workUnitRepositoryImpl) UpdateWorkUnit(id string, workUnit *models.WorkUnit) (*models.WorkUnit, error) {
	err := r.db.Model(&models.WorkUnit{}).Where("id = ?", id).Updates(workUnit).Error
	return workUnit, err
}

func (r *workUnitRepositoryImpl) DeleteWorkUnit(id string) error {
	return r.db.Delete(&models.WorkUnit{}, "id = ?", id).Error
}
