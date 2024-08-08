package models

import "gorm.io/gorm"

type Hospital struct {
	gorm.Model
	ID      uint   `json:"id" gorm:"primaryKey"`
	Name    string `json:"name" gorm:"unique;not null"`
	Email   string `json:"email" gorm:"unique;not null"`
	Phone   string `json:"phone" gorm:"unique;not null"`
	City    string `json:"city"`
	County  string `json:"county"`
	Address string `json:"address"`
}

func NewHospital(name string, email string, phone string, city string, county string, address string) *Hospital {
	return &Hospital{
		Name:    name,
		Email:   email,
		Phone:   phone,
		City:    city,
		County:  county,
		Address: address,
	}
}

func (h *Hospital) GetHospital(id int) {}
