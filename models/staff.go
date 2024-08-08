package models

import "gorm.io/gorm"

type Staff struct {
	gorm.Model
	ID           uint       `json:"-" gorm:"primaryKey"`
	FirstName    string     `json:"first_name"`
	LastName     string     `json:"last_name"`
	Email        string     `json:"email" gorm:"unique;not null"`
	Phone        string     `json:"phone" gorm:"unique;not null"`
	Password     string     `json:"password" gorm:"not null"`
	TC           string     `json:"tc" gorm:"unique;not null"`
	Privileged   bool       `json:"privileged"`
	Polyclinic   Polyclinic `json:"polyclinic" gorm:"foreignKey:PolyclinicID"`
	PolyclinicID int        `json:"polyclinic_id"`
	Title        string     `json:"title"`
	Skill        string     `json:"skill"`
}
