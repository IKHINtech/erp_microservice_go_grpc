package models

import (
	"time"

	"github.com/google/uuid"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
