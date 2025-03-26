package models

import "github.com/google/uuid"

type Permission struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name    string    `gorm:"unique;not null"` // e.g., "view_payroll"
	Service string    `gorm:"not null"`        // e.g., "hrd"
	Roles   []Role    `gorm:"many2many:role_permissions;"`
}
