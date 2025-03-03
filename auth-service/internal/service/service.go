package service

import (
	"auth-service/internal/models"
	"auth-service/internal/repository"
)

type Service interface {
	Login(user models.User) error
	Register(user models.User) error
}

type ServiceManager struct {
	repo repository.Repository
}

func NewServiceManager(repo repository.Repository) *ServiceManager {
	return &ServiceManager{repo: repo}
}

func (s *ServiceManager) Login(user models.User) error {
	return nil
}

func (s *ServiceManager) Register(user models.User) error {
	return nil
}
