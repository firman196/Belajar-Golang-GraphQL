package models

import (
	"time"

	_ "gorm.io/gorm"
)

type User struct {
	UserID    string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	Firstname string `gorm:"size:100;not null"`
	Lastname  string `gorm:"size:100"`
	Email     string `gorm:"size:100;not null;unique"`
	Password  string `gorm:"size:255,not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
