// Package: repository is an interface that defines the methods that a repository must implement to work with the cart service.
package repository

import "cart-service/internal/models"

// Repository is an interface that defines the methods that a repository must implement to work with the cart service.
type Repository interface {
	GetCart(userID int) ([]models.Product, error)
	AddProductToCart(userID, productID, quantity int) error
	DeleteProductFromCart(userID, productID int) error
	UpdateProductQuantity(userID, productID, quantity int) error
}

