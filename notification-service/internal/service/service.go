package service

import (
	"context"

	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/notification-service/internal/models"
)

type Service interface {
}

type ServiceManager struct {
	notifyChan chan models.Notification
}

func NewServiceManager(sizeChan int) *ServiceManager {
	notifyChan := make(chan models.Notification, sizeChan)
	return &ServiceManager{
		notifyChan: notifyChan,
	}
}

func (s *ServiceManager) sendNotify(notify models.Notification) error {
	// notifyChan <- notify
	return nil
}

func (s *ServiceManager) workerSendNotify(notifyChan chan models.Notification, ctx context.Context) {
	for {
		select {
		case notify := <-notifyChan:
			if err := s.sendNotify(notify); err != nil {
				// log.Printf("Failed to send notification: %s", err)
			}
		case <-ctx.Done():
			return
		}

	}
}
