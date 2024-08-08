package db

import (
	"hospital/models"
	"log"
)

func (s *Storage) AddStaff(email, phone, tc, password string, polyclinic_id int) error {
	staff_copy := models.Staff{Email: email, Phone: phone, TC: tc, Password: string(password), PolyclinicID: 2}
	result := s.DB.Create(&staff_copy)

	if result.Error != nil {
		log.Printf("[ERROR] Failed to add staff: %v", result.Error)
		return result.Error
	}

	return nil
}

func (s *Storage) PostStaff(staff models.Staff) error {
	result := s.DB.Create(&staff)
	if result.Error != nil {
		log.Printf("[ERROR] Failed to add staff: %v", result.Error)
		return result.Error
	}

	return nil
}

func (s *Storage) GetStaffWithEmail(Staff *models.Staff, email string) error {
	result := s.DB.First(&Staff, "email = ?", email)
	if result.Error != nil {
		log.Printf("[ERROR] Failed to get staff: %v", result.Error)
		return result.Error
	}

	return nil
}

func (s *Storage) GetStaffWithID(id uint) (models.Staff, error) {
	var staff models.Staff
	result := s.DB.First(&staff, "id = ?", id)
	if result.Error != nil {
		log.Printf("[ERROR] Failed to get staff: %v", result.Error)
		return models.Staff{}, result.Error
	}

	return staff, nil
}

func GetStaffWithID(Staff *models.Staff, id float64) error {
	result := HospitalStorage.DB.First(&Staff, "id = ?", id)
	if result.Error != nil {
		log.Printf("[ERROR] Failed to get staff: %v", result.Error)
		return result.Error
	}

	return nil
}

func (s *Storage) GetStaffs() ([]models.Staff, error) {
	var staffs []models.Staff
	result := s.DB.Find(&staffs)
	if result.Error != nil {
		log.Printf("[ERROR] Failed to get staffs: %v", result.Error)
		return []models.Staff{}, nil
	}

	return staffs, nil
}

func (s *Storage) GetPolyclinicStaff(polyclinic_id int) ([]models.Staff, error) {
	var staff []models.Staff
	result := s.DB.Find(&staff).Where("polyclinic_id = ?", polyclinic_id)
	if result.Error != nil {
		log.Printf("[ERROR] Failed to get staff in polyclinics %d: %v", polyclinic_id, result.Error)
		return nil, result.Error
	}

	return staff, nil
}

func (s *Storage) DeletePolyclinicStaff(polyclinic_id, staff_id int) error {
	var staff models.Staff
	result := s.DB.Delete(&staff).Where("id = ? AND polyclinic_id = ?", staff_id, polyclinic_id)
	if result.Error != nil {
		log.Printf("[ERROR] Failed to delete staff: %v", result.Error)
		return result.Error
	}

	return nil
}

func (s *Storage) DeleteStaff(id int) error {
	var staff models.Staff
	result := s.DB.Delete(&staff).Where("id = ?", id)
	if result.Error != nil {
		log.Printf("[ERROR] Failed to delete staff: %v", result.Error)
		return result.Error
	}

	return nil
}
