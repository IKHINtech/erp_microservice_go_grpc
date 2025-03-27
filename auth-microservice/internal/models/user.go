package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Email        string    `gorm:"unique;not null"`
	PasswordHash string    `gorm:"not null"`
	FullName     string    `gorm:"not null"`
	IsActive     bool      `gorm:"default:true"`
	Roles        []Role    `gorm:"many2many:user_roles;"`
	LastLogin    time.Time
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

// Hook untuk hash password sebelum create/update
func (u *User) BeforeSave(tx *gorm.DB) error {
	if u.PasswordHash != "" {
		hashedPassword, err := hashPassword(u.PasswordHash)
		if err != nil {
			return err
		}
		u.PasswordHash = hashedPassword
	}
	return nil
}

// hashPassword mengubah password menjadi hash bcrypt
func hashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// ComparePassword membandingkan password yang diinput dengan hash di database
func ComparePassword(hashedPassword, inputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	return err == nil
}
