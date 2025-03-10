package service

import (
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/catalog-service/internal/models"
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/catalog-service/internal/repository"
)

type Service interface {
	GetProducts() ([]models.Product, error)
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

func (s *ServiceManager) GetProducts() ([]models.Product, error) {
	return nil, nil
}

func (s *ServiceManager) GetProductByID(id string) (models.Product, error) {
	return models.Product{}, nil
}
