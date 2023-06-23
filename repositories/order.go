package repositories

import (
	"errors"
	"game-boost/models"
)

func CreateOrder(order models.Order) (models.Order, error) {
	if result := DB.Create(&order); result.Error != nil {
		return models.Order{}, result.Error
	}
	return order, nil
}

func GetAllOrders(userID string) ([]models.Order, error) {
	var orders []models.Order
	err := DB.Where("user_id = ?", userID).Find(&orders).Error
	if err != nil {
		return []models.Order{}, err
	}
	return orders, nil
}

func GetOrderByID(orderID string, userID string) (models.Order, error) {
	var order models.Order
	err := DB.Where("id = ? AND user_id= ?", orderID, userID).First(&order).Error
	if err != nil {
		return models.Order{}, errors.New("Order not found")
	}
	return order, nil
}

func DeleteOrder(orderID string, userID string) error {
	order := models.Order{}
	result := DB.Unscoped().Where("id = ? AND user_id= ?", orderID, userID).Delete(&order)
	if result.RowsAffected == 0 {
		return errors.New("Failed to delete order")
	}

	return nil
}

func EditOrder(order *models.Order, orderID string, userID string) error {
	result := DB.Model(&order).Where("id = ? AND user_id= ?", orderID, userID).Updates(order)
	if result.RowsAffected == 0 {
		return errors.New("failed to update order")
	}
	return nil
}
