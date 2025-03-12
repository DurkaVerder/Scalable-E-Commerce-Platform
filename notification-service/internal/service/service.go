package service

import "github.com/DurkaVerder/Scalable-E-Commerce-Platform/notification-service/internal/models"

type Service interface {
	SendNotification(notify models.Notification) error
}

type ServiceManager struct {
}

func NewServiceManager() *ServiceManager {
	return &ServiceManager{}
}

func (sm *ServiceManager) SendNotification(notify models.Notification) error {
	// Here we would send the notification to the user
	// For now, we just return nil
	return nil
}
