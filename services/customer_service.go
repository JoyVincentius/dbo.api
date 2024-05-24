package services

import (
	"dbo.api/database"
	"dbo.api/models"
)

func GetAllCustomers() ([]models.Customer, error) {
	var customers []models.Customer
	if err := database.DB.Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}

func GetCustomerByID(id uint) (models.Customer, error) {
	var customer models.Customer
	if err := database.DB.First(&customer, id).Error; err != nil {
		return customer, err
	}
	return customer, nil
}

func CreateCustomer(customer models.Customer) (models.Customer, error) {
	if err := database.DB.Create(&customer).Error; err != nil {
		return customer, err
	}
	return customer, nil
}

func UpdateCustomer(id uint, updatedCustomer models.Customer) (models.Customer, error) {
	var customer models.Customer
	if err := database.DB.First(&customer, id).Error; err != nil {
		return customer, err
	}

	customer.Name = updatedCustomer.Name
	customer.Email = updatedCustomer.Email

	if err := database.DB.Save(&customer).Error; err != nil {
		return customer, err
	}
	return customer, nil
}

func DeleteCustomer(id uint) error {
	if err := database.DB.Delete(&models.Customer{}, id).Error; err != nil {
		return err
	}
	return nil
}

func SearchCustomers(query string, page int, limit int) ([]models.Customer, error) {
	var customers []models.Customer
	offset := (page - 1) * limit

	if err := database.DB.Where("name LIKE ? OR email LIKE ?", "%"+query+"%", "%"+query+"%").
		Offset(offset).
		Limit(limit).
		Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}
