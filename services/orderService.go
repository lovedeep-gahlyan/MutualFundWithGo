package services

import (
    "mutualfund/models"
	"mutualfund/repositories"
)

type OrderService struct {
    orderRepository *repositories.OrderRepository
}

func NewOrderService(orderRepository *repositories.OrderRepository) *OrderService {
    return &OrderService{orderRepository: orderRepository}
}

// CreateOrder creates a new order using the order repository
func (s OrderService) CreateOrder(order *models.Order) (*models.Order, *models.ResponseError) {
    return s.orderRepository.CreateOrder(order)
}

// GetOrdersByUserID retrieves all orders for a specific user
func (s OrderService) GetOrdersByUserID(userID int64) ([]models.Order, *models.ResponseError) {
    return s.orderRepository.GetOrdersByUserID(userID)
}
