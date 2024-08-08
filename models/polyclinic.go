package models

import "gorm.io/gorm"

type Polyclinic struct {
	gorm.Model
	ID         int      `gorm:"primaryKey"`
	Name       string   `json:"name"`
	City       string   `json:"city"`
	County     string   `json:"county"`
	Address    string   `json:"address"`
	TotalStaff uint32   `json:"total_staff"`
	Hospital   Hospital `json:"hospital" gorm:"foreignKey:HospitalID"`
	HospitalID int      `json:"hospital_id"`
}
