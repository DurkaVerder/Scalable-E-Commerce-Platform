package repository

import "auth-service/internal/models"

type Repository interface {
	GetUser(email string) (models.User, error)
	CreateUser(models.User) error
}
