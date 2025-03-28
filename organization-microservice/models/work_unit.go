package models

import "github.com/google/uuid"

// WorkUnit model
type WorkUnit struct {
	BaseModel
	DepartmentID uuid.UUID  `gorm:"type:uuid"`
	Department   Department `gorm:"foreignKey:DepartmentID;references:ID;onUpdate:CASCADE;onDelete:CASCADE"`
	Name         string
	Description  string
}
