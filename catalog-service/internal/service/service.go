package service

import (
	"strconv"

	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/catalog-service/internal/models"
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/catalog-service/internal/repository"
	elk "github.com/DurkaVerder/elk-send-logs/elk"
)

type Service interface {
	GetProducts(category string) ([]models.Product, error)
	GetProductByID(id string) (models.Product, error)
}

type ServiceManager struct {
	repo repository.Repository
}

func NewServiceManager(repo repository.Repository) *ServiceManager {
	return &ServiceManager{
		repo: repo,
	}
}

func (s *ServiceManager) GetProducts(category string) ([]models.Product, error) {
	var err error
	products := []models.Product{}

	if category != "" {
		products, err = s.repo.GetProductsByCategory(category)
		if err != nil {
			elk.Log.SendMsg(
				elk.LogMessage{
					Level:   'E',
					Message: "Error getting products by category",
					Fields: map[string]interface{}{
						"method":   "GetProducts",
						"category": category,
						"error":    err.Error(),
					},
				})
			return nil, err
		}
	} else {
		products, err = s.repo.GetAllProducts()
		if err != nil {
			elk.Log.SendMsg(
				elk.LogMessage{
					Level:   'E',
					Message: "Error getting all products",
					Fields: map[string]interface{}{
						"method": "GetProducts",
						"error":  err.Error(),
					},
				})
			return nil, err
		}
	}

	return products, nil
}

func (s *ServiceManager) GetProductByID(id string) (models.Product, error) {
	productId, err := strconv.Atoi(id)
	if err != nil {
		elk.Log.SendMsg(
			elk.LogMessage{
				Level:   'E',
				Message: "Error converting id to int",
				Fields: map[string]interface{}{
					"method": "GetProductByID",
					"id":     id,
					"error":  err.Error(),
				},
			})
		return models.Product{}, err
	}

	product, err := s.repo.GetProductById(productId)
	if err != nil {
		elk.Log.SendMsg(
			elk.LogMessage{
				Level:   'E',
				Message: "Error getting product by id",
				Fields: map[string]interface{}{
					"method": "GetProductByID",
					"id":     id,
					"error":  err.Error(),
				},
			})
		return models.Product{}, err
	}

	return product, nil
}
