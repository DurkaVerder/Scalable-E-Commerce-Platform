package repository

import "cart-service/internal/models"

type Repository interface {
	GetCart(userID int) ([]models.Product, error)
	AddProductToCart(userID, productID, quantity int) error
	DeleteProductFromCart(userID, productID int) error
	UpdateProductQuantity(userID, productID, quantity int) error
}
