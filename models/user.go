package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null;type:varchar(255)" json:"email"`
	Password string `gorm:"not null;type:varchar(255)" json:"password"`
}