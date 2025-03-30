package repository

import (
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/internal/models"
	"gorm.io/gorm"
)

type OrganizationRepository interface {
	CreateOrganization(organization *models.Organization) (*models.Organization, error)
	GetOrganization(id string) (*models.Organization, error)
	GetListOrganization(parentOrganizationID *string) ([]*models.Organization, error)
	UpdateOrganization(id string, organization *models.Organization) (*models.Organization, error)
	DeleteOrganization(id string) error
}

type organizationRepositoryImpl struct {
	db *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) OrganizationRepository {
	return &organizationRepositoryImpl{
		db: db,
	}
}

func (r *organizationRepositoryImpl) CreateOrganization(organization *models.Organization) (*models.Organization, error) {
	if err := r.db.Create(organization).Error; err != nil {
		return nil, err
	}
	return organization, nil
}

func (r *organizationRepositoryImpl) GetOrganization(id string) (*models.Organization, error) {
	var organization models.Organization
	err := r.db.Where("id = ?", id).First(&organization).Error
	return &organization, err
}

func (r *organizationRepositoryImpl) GetListOrganization(parentOrganizationID *string) ([]*models.Organization, error) {
	var organizations []*models.Organization
	query := r.db
	if parentOrganizationID != nil {
		query = query.Where("parent_organization_id = ?", parentOrganizationID)
	}
	err := query.Find(&organizations).Error
	return organizations, err
}

func (r *organizationRepositoryImpl) UpdateOrganization(id string, organization *models.Organization) (*models.Organization, error) {
	err := r.db.Model(&models.Organization{}).Where("id = ?", id).Updates(organization).Error
	return organization, err
}

func (r *organizationRepositoryImpl) DeleteOrganization(id string) error {
	return r.db.Delete(&models.Organization{}, "id = ?", id).Error
}
