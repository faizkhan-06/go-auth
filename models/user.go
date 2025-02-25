package models

import "time"

type User struct {
	ID        uint	`gorm:"primaryKey"`
	Email     string `gorm:"unique"`
	Password  string	
	CreatedAt time.Time `gorm:"autoIncrement"`
	UpdatedAt time.Time	`gorm:"autoIncrement"`
	DeletedAt time.Time	`gorm:"autoIncrement"`
}