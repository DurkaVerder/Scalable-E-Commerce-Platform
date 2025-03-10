package repository

import "github.com/DurkaVerder/Scalable-E-Commerce-Platform/catalog-service/internal/models"

type Repository interface {
	GetAllProducts() ([]models.Product, error)
	GetProductById(id int) (models.Product, error)
	GetProductsByCategory(category string) ([]models.Product, error)
}
