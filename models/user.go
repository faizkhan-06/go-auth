package models

import "time"

type User struct {
	ID        uint	`gorm:"primaryKey"`
	Email     string `gorm:"unique;not null;type:varchar(255)"`
	Password  string	`gorm:"not null;type:varchar(255)"`
	CreatedAt time.Time `gorm:"autoIncrement"`
	UpdatedAt time.Time	`gorm:"autoIncrement"`
	DeletedAt time.Time	`gorm:"autoIncrement"`
}