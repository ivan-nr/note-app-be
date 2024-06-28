package model

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Title  string `json:"title" gorm:"not null"`
	Note   string `json:"note"`
	UserId int    `gorm:"foreignKey" json:"user_id"`
}
