package db

import (
	"hospital/models"
	"log"
)

func (s *Storage) CreateHospital(h models.Hospital) {
	s.DB.Create(&models.Hospital{Name: h.Name, Email: h.Email, Phone: h.Phone, City: h.City, County: h.County, Address: h.Address})
}

// Get hospital with given ID
func (s *Storage) GetHospital(pk int) (models.Hospital, error) {
	var hospital models.Hospital
	result := s.DB.First(&hospital, pk)
	if result.Error != nil {
		log.Printf("[ERROR] Failed to get hospital: %v", result.Error)
		return models.Hospital{}, result.Error
	}

	return hospital, nil
}

func (s *Storage) GetHospitals() ([]models.Hospital, error) {
	var hospitals []models.Hospital
	result := s.DB.Find(&hospitals)
	if result.Error != nil {
		log.Printf("[ERROR] Failed to get hospitals: %v", result.Error)
		return []models.Hospital{}, nil
	}

	return hospitals, nil
}

// Get hospital from database with given name
func (s *Storage) GetHospitalWithName(name string) (models.Hospital, error) {
	var hospital models.Hospital
	result := s.DB.First(&hospital, "name = ?", name)
	if result.Error != nil {
		log.Printf("[ERROR] Failed to get hospital with name: %v", result.Error)
		return models.Hospital{}, result.Error
	}

	return hospital, nil
}

// Get hospital with given ID
func GetHospitalWithID(pk int) (*models.Hospital, error) {
	var hospital models.Hospital
	result := HospitalStorage.DB.First(&hospital, pk)
	if result.Error != nil {
		log.Printf("[ERROR] Failed to get hospital: %v", result.Error)
		return nil, result.Error
	}

	return &hospital, nil
}

// Add hospital to database with given hospital model
func (s *Storage) AddHospital(hospital *models.Hospital) error {
	result := s.DB.Create(&hospital)
	if result.Error != nil {
		log.Printf("[ERROR] Failed to add hospital: %v", result.Error)
		return result.Error
	}

	return nil
}

// Delete hospital with given ID
func (s *Storage) DeleteHospital(id int) error {
	hospital, err := s.GetHospital(id)
	if err != nil {
		log.Printf("[ERROR] Failed to get hospital with ID: %d", id)
		return err
	}

	result := s.DB.Delete(&hospital)
	if result.Error != nil {
		log.Printf("[ERROR] Failed to delete hospital: %v", result.Error)
		return err
	}

	return nil
}
