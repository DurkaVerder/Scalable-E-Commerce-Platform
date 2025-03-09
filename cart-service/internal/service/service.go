package service

import (
	"cart-service/internal/models"
	"cart-service/internal/repository"
)

type Service interface {
	GetCart(userID int) ([]models.Product, error)
	AddProductToCart(userID, productID, quantity int) error
	RemoveProductFromCart(userID, productID int) error
	UpdateProductQuantity(userID, productID, quantity int) error
}

type ServiceManager struct {
	repo repository.Repository
}

func NewServiceManager(repo repository.Repository) *ServiceManager {
	return &ServiceManager{repo: repo}
}

func (s *ServiceManager) GetCart(userID int) ([]models.Product, error) {
	return nil, nil
}

func (s *ServiceManager) AddProductToCart(userID, productID, quantity int) error {
	return nil
}

func (s *ServiceManager) RemoveProductFromCart(userID, productID int) error {
	return nil
}

func (s *ServiceManager) UpdateProductQuantity(userID, productID, quantity int) error {
	return nil
}
