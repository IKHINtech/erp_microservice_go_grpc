package models

import (
	"time"

	"github.com/google/uuid"
)

type Position struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Title       string    `gorm:"not null;unique"`
	JobCode     string    `gorm:"not null"`
	PayGrade    string    `gorm:"not null"`
	Description string    `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
