package models

import "gorm.io/gorm"

type Skill struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey"`
	Name string `json:"name" gorm:"unique"`
}
