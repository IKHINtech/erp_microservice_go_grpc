package models

// Organization model
type Organization struct {
	BaseModel
	Name                 string
	Description          string
	Type                 string
	ParentOrganizationID *string       `gorm:"type:uuid"`
	ParentOrganization   *Organization `gorm:"foreignKey:ParentOrganizationID;references:ID"`
}
