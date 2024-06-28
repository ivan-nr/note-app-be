package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null"`
	Username string `gorm:"uniqueIndex;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
}
