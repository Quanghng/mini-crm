package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Name  string `json:"name" gorm:"type:varchar(100);not null"`
	Email string `json:"email" gorm:"type:varchar(100);not null"`
}
