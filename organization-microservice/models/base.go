package models

import (
	"time"
)

type BaseModel struct {
	ID        string `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
