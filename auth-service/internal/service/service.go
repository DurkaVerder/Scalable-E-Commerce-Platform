package service

import (
	"auth-service/internal/models"
	"auth-service/internal/repository"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(user models.User) (string, error)
	Register(user models.User) error
}

type ServiceManager struct {
	repo repository.Repository
}

func NewServiceManager(repo repository.Repository) *ServiceManager {
	return &ServiceManager{repo: repo}
}

func (s *ServiceManager) Login(user models.User) (string, error) {
	storedUser, err := s.repo.GetUser(user.Email)
	if err != nil {
		log.Printf("Failed to get user: %v", err)
		return "", fmt.Errorf("not found")
	}

	if err := comparePasswords(storedUser.Password, user.Password); err != nil {
		log.Printf("Failed to compare passwords: %v", err)
		return "error", err
	}

	token, err := s.generateJWT(storedUser.ID)
	if err != nil {
		log.Printf("Failed to generate JWT: %v", err)
		return "", err
	}

	return token, nil
}

func (s *ServiceManager) Register(user models.User) error {
	if !s.validPassword(user.Password) {
		return fmt.Errorf("invalid password")
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	if err := s.repo.CreateUser(user); err != nil {
		log.Printf("Failed to create user: %v", err)
		return err
	}

	return nil
}

func (s *ServiceManager) validPassword(password string) bool {
	return len(password) >= 6
}

func hashPassword(password string) (string, error) {
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return "", err
	}

	return string(bcryptPassword), nil
}

func comparePasswords(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
