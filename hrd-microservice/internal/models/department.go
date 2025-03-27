package models

import (
	"time"

	"github.com/google/uuid"
)

type Department struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name      string     `gorm:"not null;unique"`
	Code      string     `gorm:"not null;unique"`
	ManagerID *uuid.UUID `gorm:"type:uuid"`
	Manager   *Employee  `gorm:"foreignKey:ManagerID"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
}
