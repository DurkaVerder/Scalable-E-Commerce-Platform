package service

import (
	"database/sql"
	"fmt"

	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/auth-service/internal/models"
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/auth-service/internal/repository"
	elk "github.com/DurkaVerder/Scalable-E-Commerce-Platform/auth-service/pkg/logs"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(user models.User) (string, error)
	Register(user models.User) error
	ValidateJWT(token string) error
	GetUserIdFromToken(token string) (int, error)
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
		elk.Log.Error("Error getting user", map[string]interface{}{
			"method": "Login",
			"action": "GetUser",
			"error":  err.Error(),
		})

		return "", fmt.Errorf("not found")
	}

	if err := comparePasswords(storedUser.Password, user.Password); err != nil {
		elk.Log.Error("Error comparing passwords", map[string]interface{}{
			"method": "Login",
			"action": "ComparePasswords",
			"error":  err.Error(),
		})
		return "error", err
	}

	token, err := s.generateJWT(storedUser.ID)
	if err != nil {
		elk.Log.Error("Error generating JWT", map[string]interface{}{
			"method": "Login",
			"action": "GenerateJWT",
			"error":  err.Error(),
		})
		return "", err
	}

	return token, nil
}

func (s *ServiceManager) Register(user models.User) error {
	_, err := s.repo.GetUser(user.Email)
	if err != sql.ErrNoRows {
		elk.Log.Error("User already exists", map[string]interface{}{
			"method": "Register",
			"action": "GetUser",
			"error":  err.Error(),
		})
		return fmt.Errorf("user already exists")
	}

	if !s.validPassword(user.Password) {
		elk.Log.Error("Invalid password", map[string]interface{}{
			"method": "Register",
			"action": "ValidPassword",
		})
		return fmt.Errorf("invalid password")
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		elk.Log.Error("Error hashing password", map[string]interface{}{
			"method": "Register",
			"action": "HashPassword",
			"error":  err.Error(),
		})
		return err
	}

	user.Password = hashedPassword

	if err := s.repo.CreateUser(user); err != nil {
		elk.Log.Error("Error creating user", map[string]interface{}{
			"method": "Register",
			"action": "CreateUser",
			"error":  err.Error(),
		})
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
		return "", err
	}

	return string(bcryptPassword), nil
}

func comparePasswords(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
