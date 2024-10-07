package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `gorm:"size:255;not null"`
	Email string `gorm:"size:255;not null"`
	// BornDate time.
}
