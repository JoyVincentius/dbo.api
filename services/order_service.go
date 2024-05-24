package services

import (
	"dbo.api/database"
	"dbo.api/models"
)

func GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	if err := database.DB.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func GetOrderByID(id uint) (models.Order, error) {
	var order models.Order
	if err := database.DB.First(&order, id).Error; err != nil {
		return order, err
	}
	return order, nil
}

func CreateOrder(order models.Order) (models.Order, error) {
	if err := database.DB.Create(&order).Error; err != nil {
		return order, err
	}
	return order, nil
}

func UpdateOrder(id uint, orderData models.Order) (models.Order, error) {
	var order models.Order
	if err := database.DB.First(&order, id).Error; err != nil {
		return order, err
	}

	order.CustomerID = orderData.CustomerID
	order.Item = orderData.Item
	order.Quantity = orderData.Quantity

	if err := database.DB.Save(&order).Error; err != nil {
		return order, err
	}
	return order, nil
}

func DeleteOrder(id uint) error {
	if err := database.DB.Delete(&models.Order{}, id).Error; err != nil {
		return err
	}
	return nil
}

func SearchOrders(query string, page int, limit int) ([]models.Order, error) {
	var orders []models.Order
	offset := (page - 1) * limit

	if err := database.DB.Where("item LIKE ?", "%"+query+"%").
		Offset(offset).
		Limit(limit).
		Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
