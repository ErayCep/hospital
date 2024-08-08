package db

import (
	"hospital/models"
	"log"
)

func (s *Storage) GetPolyclinics() ([]models.Polyclinic, error) {
	var polyclinics []models.Polyclinic
	result := s.DB.Find(&polyclinics)
	if result.Error != nil {
		log.Printf("[ERROR] Failed to get polyclinics: %v", result.Error)
		return []models.Polyclinic{}, result.Error
	}

	return polyclinics, nil
}

func (s *Storage) GetPolyclinic(hospital_id, polyclinic_id int) (models.Polyclinic, error) {
	var polyclinic models.Polyclinic
	result := s.DB.Where("id = ? AND hospital_id = ?", polyclinic_id, hospital_id).First(&polyclinic)
	if result.Error != nil {
		log.Printf("[ERROR] Failed to get polyclinic: %v", result.Error)
		return models.Polyclinic{}, result.Error
	}

	return polyclinic, nil
}

func (s *Storage) GetPolyclinic2(polyclinic_id int) (models.Polyclinic, error) {
	var polyclinic models.Polyclinic
	result := s.DB.Where("id = ?", polyclinic_id).First(&polyclinic)
	if result.Error != nil {
		log.Printf("[ERROR] Failed to get polyclinic: %v", result.Error)
		return models.Polyclinic{}, result.Error
	}

	return polyclinic, nil
}

func (s *Storage) CreatePolyclinic(city, county, address string, total_staff uint32, hospital_id int) error {
	polyclinic := models.Polyclinic{City: city, County: county, Address: address, TotalStaff: total_staff, HospitalID: hospital_id}
	result := s.DB.Create(&polyclinic)
	if result.Error != nil {
		log.Printf("[ERROR] Failed to create polyclinic: %v", result.Error)
		return result.Error
	}

	return nil
}

func (s *Storage) AddPolyclinic(polyclinic models.Polyclinic) error {
	result := s.DB.Create(&polyclinic)
	if result.Error != nil {
		log.Printf("[ERROR] Failed to create polyclinic: %v", result.Error)
		return result.Error
	}

	return nil
}

func (s *Storage) DeletePolyclinic(id int) error {
	var polyclinic models.Polyclinic
	result := s.DB.Delete(&polyclinic, id)
	if result.Error != nil {
		log.Printf("[ERROR] Failed to delete polyclinic: %v", result.Error)
		return result.Error
	}

	return nil
}
