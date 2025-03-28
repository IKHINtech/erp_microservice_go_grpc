package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PositionHierarchy model
type PositionHierarchy struct {
	gorm.Model
	SuperiorPositionID     *uuid.UUID `gorm:"type:uuid;null"`
	SubordinatedPositionID *uuid.UUID `gorm:"type:uuid;null"`
	SuperiorPosition       *Position  `gorm:"foreignKey:SuperiorPositionID;references:ID"`
	SubordinatedPosition   *Position  `gorm:"foreignKey:SubordinatedPositionID;references:ID"`
}
