package models

import (
	"gorm.io/gorm"
)

// PositionHierarchy model
type PositionHierarchy struct {
	gorm.Model
	SuperiorPositionID     *string   `gorm:"type:uuid;null"`
	SubordinatedPositionID *string   `gorm:"type:uuid;null"`
	SuperiorPosition       *Position `gorm:"foreignKey:SuperiorPositionID;references:ID"`
	SubordinatedPosition   *Position `gorm:"foreignKey:SubordinatedPositionID;references:ID"`
}
