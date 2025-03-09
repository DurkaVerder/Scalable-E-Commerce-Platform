package repository

import "github.com/DurkaVerder/Scalable-E-Commerce-Platform/auth-service/internal/models"

type Repository interface {
	GetUser(email string) (models.User, error)
	CreateUser(models.User) error
}
