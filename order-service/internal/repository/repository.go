package repository

import "github.com/DurkaVerder/Scalable-E-Commerce-Platform/order-service/internal/models"

type Repository interface {
	CreateOrder(userID int, amount float64, products []models.Product) error
	GetOrder(orderID int) (models.Order, error)
	GetOrders(userID int) ([]models.Order, error)
	GetOrderProducts(orderID int) ([]models.OrderProduct, error)
	UpdateOrder(orderID int, status string) error
	Close()
}
