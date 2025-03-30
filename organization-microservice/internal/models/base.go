package models

import (
	"time"
)

type BaseModel struct {
	ID        string `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
