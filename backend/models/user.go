package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null;size:100"`
	Password string `gorm:"not null;size:255"`
}
