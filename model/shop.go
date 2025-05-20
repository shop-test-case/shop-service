package model

import "gorm.io/gorm"

type Shop struct {
	gorm.Model
	Name        string `json:"name" gorm:"unique"`
	Description string `json:"description"`
}
