package models

// WorkUnit model
type WorkUnit struct {
	BaseModel
	DepartmentID string     `gorm:"type:uuid"`
	Department   Department `gorm:"foreignKey:DepartmentID;references:ID;onUpdate:CASCADE;onDelete:CASCADE"`
	Name         string
	Description  string
}
