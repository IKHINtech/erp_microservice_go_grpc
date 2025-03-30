package models

import (
	"github.com/google/uuid"
)

// Organization model
type Organization struct {
	BaseModel
	Name                 string
	Description          string
	Type                 string
	ParentOrganizationID *string       `gorm:"type:uuid"`
	ParentOrganization   *Organization `gorm:"foreignKey:ParentOrganizationID;references:ID"`
}

// CreateOrganizationRequest model
type CreateOrganizationRequest struct {
	Name                 string
	Description          string
	Type                 string
	ParentOrganizationID string `gorm:"type:uuid"`
}

// GetOrganizationRequest model
type GetOrganizationRequest struct {
	ID uuid.UUID `gorm:"type:uuid"`
}

// GetOrganizationResponse model
type GetOrganizationResponse struct {
	Organization Organization
}

// ListOrganizationRequest model
type ListOrganizationRequest struct {
	Type                 string
	ParentOrganizationID uuid.UUID `gorm:"type:uuid"`
}

// ListOrganizationResponse model
type ListOrganizationResponse struct {
	Organizations []Organization
}

// UpdateOrganizationRequest model
type UpdateOrganizationRequest struct {
	ID                   uuid.UUID `gorm:"type:uuid"`
	Name                 string
	Description          string
	Type                 string
	ParentOrganizationID uuid.UUID `gorm:"type:uuid"`
}

// DeleteOrganizationRequest model
type DeleteOrganizationRequest struct {
	ID uuid.UUID `gorm:"type:uuid"`
}

// CreateDepartmentRequest model
type CreateDepartmentRequest struct {
	Name        string
	Description string
}

// GetDepartmentRequest model
type GetDepartmentRequest struct {
	ID uuid.UUID `gorm:"type:uuid"`
}

// GetDepartmentResponse model
type GetDepartmentResponse struct {
	Department Department
}

// ListDepartmentRequest model
type ListDepartmentRequest struct{}

// ListDepartmentResponse model
type ListDepartmentResponse struct {
	Departments []Department
}

// UpdateDepartmentRequest model
type UpdateDepartmentRequest struct {
	ID          uuid.UUID `gorm:"type:uuid"`
	Name        string
	Description string
}

// DeleteDepartmentRequest model
type DeleteDepartmentRequest struct {
	ID uuid.UUID `gorm:"type:uuid"`
}

// CreateWorkUnitRequest model
type CreateWorkUnitRequest struct {
	DepartmentID uuid.UUID `gorm:"type:uuid"`
	Name         string
	Description  string
}

// GetWorkUnitRequest model
type GetWorkUnitRequest struct {
	ID uuid.UUID `gorm:"type:uuid"`
}

// GetWorkUnitResponse model
type GetWorkUnitResponse struct {
	WorkUnit WorkUnit
}

// ListWorkUnitRequest model
type ListWorkUnitRequest struct {
	DepartmentID uuid.UUID `gorm:"type:uuid"`
}

// ListWorkUnitResponse model
type ListWorkUnitResponse struct {
	WorkUnits []WorkUnit
}

// UpdateWorkUnitRequest model
type UpdateWorkUnitRequest struct {
	ID           uuid.UUID `gorm:"type:uuid"`
	DepartmentID uuid.UUID `gorm:"type:uuid"`
	Name         string
	Description  string
}

// DeleteWorkUnitRequest model
type DeleteWorkUnitRequest struct {
	ID uuid.UUID `gorm:"type:uuid"`
}

// CreatePositionRequest model
type CreatePositionRequest struct {
	Name        string
	Description string
}

// GetPositionRequest model
type GetPositionRequest struct {
	ID uuid.UUID `gorm:"type:uuid"`
}

// GetPositionResponse model
type GetPositionResponse struct {
	Position Position
}

// ListPositionRequest model
type ListPositionRequest struct{}

// ListPositionResponse model
type ListPositionResponse struct {
	Positions []Position
}

// UpdatePositionRequest model
type UpdatePositionRequest struct {
	ID          uuid.UUID `gorm:"type:uuid"`
	Name        string
	Description string
}

// DeletePositionRequest model
type DeletePositionRequest struct {
	ID uuid.UUID `gorm:"type:uuid"`
}

// CreatePositionHierarchyRequest model
type CreatePositionHierarchyRequest struct {
	SuperiorPositionID     uuid.UUID `gorm:"type:uuid"`
	SubordinatedPositionID uuid.UUID `gorm:"type:uuid"`
}

// GetPositionHierarchyBySuperiorRequest model
type GetPositionHierarchyBySuperiorRequest struct {
	SuperiorPositionID uuid.UUID `gorm:"type:uuid"`
}

// GetPositionHierarchyBySubordinatedRequest model
type GetPositionHierarchyBySubordinatedRequest struct {
	SubordinatedPositionID uuid.UUID `gorm:"type:uuid"`
}

// ListPositionHierarchyResponse model
type ListPositionHierarchyResponse struct {
	PositionHierarchy []PositionHierarchy
}

// DeletePositionHierarchyRequest model
type DeletePositionHierarchyRequest struct {
	ID uuid.UUID `gorm:"type:uuid"`
}

// GetPositionHierarchyResponse model
type GetPositionHierarchyResponse struct {
	PositionHierarchy []PositionHierarchy
}
