package service

import (
	"strconv"

	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/order-service/internal/models"
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/order-service/internal/repository"
	elk "github.com/DurkaVerder/Scalable-E-Commerce-Platform/order-service/pkg/logs"
)

type Service interface {
	CreateOrder(userId int, products []models.Product) error
	GetOrders(userId int) ([]models.Order, error)
	GetOrder(orderId int) ([]models.OrderProduct, error)
	UpdateOrder(orderId int, status string) error
	ConvertStringToInt(num string) (int, error)
	Close()
}

type ServiceManager struct {
	repo repository.Repository
}

func NewServiceManager(repo repository.Repository) *ServiceManager {
	return &ServiceManager{
		repo: repo,
	}
}

func (s *ServiceManager) Close() {
	s.repo.Close()
}

func (s *ServiceManager) CreateOrder(userId int, products []models.Product) error {
	amount := s.sumProducts(products)

	if err := s.repo.CreateOrder(userId, amount, products); err != nil {
		elk.Log.Error("Error creating order", map[string]interface{}{
			"method": "CreateOrder",
			"action": "CreateOrder",
			"error":  err,
		})
		return err
	}
	return nil
}

func (s *ServiceManager) sumProducts(products []models.Product) float64 {
	var totalSum float64
	for _, product := range products {
		sum := product.Price * float64(product.Quantity)
		totalSum += sum
	}
	return totalSum
}

func (s *ServiceManager) GetOrders(userId int) ([]models.Order, error) {
	orders, err := s.repo.GetOrders(userId)
	if err != nil {
		elk.Log.Error("Error getting orders", map[string]interface{}{
			"method": "GetOrders",
			"action": "GetOrders",
			"error":  err,
		})
		return nil, err
	}
	return orders, nil
}

func (s *ServiceManager) GetOrder(orderId int) ([]models.OrderProduct, error) {

	order, err := s.repo.GetOrder(orderId)
	if err != nil {
		elk.Log.Error("Error getting order", map[string]interface{}{
			"method":   "GetOrder",
			"action":   "GetOrder",
			"order_id": orderId,
			"error":    err,
		})
		return nil, err
	}

	orderProduct, err := s.repo.GetOrderProducts(orderId)
	if err != nil {
		elk.Log.Error("Error getting order items", map[string]interface{}{
			"method":   "GetOrderItems",
			"action":   "GetOrderItems",
			"order_id": orderId,
			"error":    err,
		})
		return nil, err
	}

	order.Products = orderProduct
	return order.Products, nil
}

func (s *ServiceManager) UpdateOrder(orderId int, status string) error {
	if err := s.repo.UpdateOrder(orderId, status); err != nil {
		elk.Log.Error("Error updating order", map[string]interface{}{
			"method":   "UpdateOrder",
			"action":   "UpdateOrder",
			"order_id": orderId,
			"error":    err,
		})
		return err
	}
	return nil
}

func (s *ServiceManager) ConvertStringToInt(num string) (int, error) {
	resInt, err := strconv.Atoi(num)
	if err != nil {
		elk.Log.Error("Error converting string to int", map[string]interface{}{
			"method": "convertStringToInt",
			"action": "convertStringToInt",
			"error":  err,
		})
		return 0, err
	}
	return resInt, nil
}
