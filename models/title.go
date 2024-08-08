package models

import "gorm.io/gorm"

type Title struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey"`
	Value string `json:"value" gorm:"unique"`
}
