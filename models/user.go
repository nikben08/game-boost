package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();"`
	Password string    `gorm:"hash"`
	Email    string    `gorm:"email"`
	Name     string    `gorm:"name"`
}
