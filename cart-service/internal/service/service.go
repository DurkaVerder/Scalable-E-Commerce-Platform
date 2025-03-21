// Package: service provides the business logic for the cart service.
package service

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/cart-service/internal/models"
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/cart-service/internal/repository"
	elk "github.com/DurkaVerder/Scalable-E-Commerce-Platform/cart-service/pkg/logs"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetCart(userID int) ([]models.Product, error)
	AddProductToCart(userID, productID, quantity int) error
	RemoveProductFromCart(userID, productID int) error
	UpdateProductQuantity(userID, productID, quantity int) error
	GetUserID(ctx *gin.Context) (int, error)
}

// ServiceManager is a struct that implements the Service interface.
type ServiceManager struct {
	repo repository.Repository
}

// NewServiceManager creates a new service manager.
func NewServiceManager(repo repository.Repository) *ServiceManager {
	return &ServiceManager{repo: repo}
}

// GetCart returns the cart of a user.
func (s *ServiceManager) GetCart(userID int) ([]models.Product, error) {
	products, err := s.repo.GetCart(userID)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		elk.Log.SendMsg(elk.LogMessage{
			Level:   'E',
			Message: "Error getting cart from repository",
			Fields: map[string]interface{}{
				"method":  "GetCart",
				"action":  "getting cart from repository",
				"error":   err,
				"user_id": userID,
			},
		})

		return nil, err
	}

	return products, nil
}

// AddProductToCart adds a product to the cart.
func (s *ServiceManager) AddProductToCart(userID, productID, quantity int) error {
	if err := s.repo.AddProductToCart(userID, productID, quantity); err != nil {
		elk.Log.SendMsg(elk.LogMessage{
			Level:   'E',
			Message: "Error adding product to cart in repository",
			Fields: map[string]interface{}{
				"method":    "AddProductToCart",
				"action":    "adding product to cart in repository",
				"error":     err,
				"user_id":   userID,
				"productID": productID,
				"quantity":  quantity,
			},
		})

		return err
	}

	return nil
}

// RemoveProductFromCart removes a product from the cart.
func (s *ServiceManager) RemoveProductFromCart(userID, productID int) error {
	if err := s.repo.DeleteProductFromCart(userID, productID); err != nil {
		elk.Log.SendMsg(elk.LogMessage{
			Level:   'E',
			Message: "Error removing product from cart in repository",
			Fields: map[string]interface{}{
				"method":    "RemoveProductFromCart",
				"action":    "removing product from cart in repository",
				"error":     err,
				"user_id":   userID,
				"productID": productID,
			},
		})
		return err
	}

	return nil
}

// UpdateProductQuantity updates the quantity of a product in the cart.
// If the quantity is 0, the product is removed from the cart.
func (s *ServiceManager) UpdateProductQuantity(userID, productID, quantity int) error {
	if quantity == 0 {
		return s.RemoveProductFromCart(userID, productID)
	}

	if err := s.repo.UpdateProductQuantity(userID, productID, quantity); err != nil {
		elk.Log.SendMsg(elk.LogMessage{
			Level:   'E',
			Message: "Error updating product quantity in repository",
			Fields: map[string]interface{}{
				"method":    "UpdateProductQuantity",
				"action":    "updating product quantity in repository",
				"error":     err,
				"user_id":   userID,
				"productID": productID,
				"quantity":  quantity,
			},
		})
		return err
	}

	return nil
}

// GetUserID returns the user ID from the request context.
func (s *ServiceManager) GetUserID(ctx *gin.Context) (int, error) {
	userID := ctx.GetHeader("user_id")
	if userID == "" {
		return -1, fmt.Errorf("user_id header is missing")
	}

	id, err := strconv.Atoi(userID)
	if err != nil {
		return -1, err
	}
	return id, nil
}
