package repository

import (
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/payment-service/internal/models"
)

type Repository interface {
	AddPayment(order models.Order, paymentIntendId string) error
	Close()
}
