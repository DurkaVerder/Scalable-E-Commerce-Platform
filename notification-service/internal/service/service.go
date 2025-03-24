package service

import (
	"context"
	"os"

	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/notification-service/internal/models"
	elk "github.com/DurkaVerder/elk-send-logs/elk"
	"gopkg.in/gomail.v2"
)

type Service interface {
	InputNotify(notify models.Notification)
	Start(ctx context.Context, countWorker int)
	Close()
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

	msg := gomail.NewMessage()
	from := os.Getenv("MAIL")
	msg.SetHeader("From", from)
	msg.SetHeader("To", notify.Email)
	msg.SetHeader("Subject", notify.Subject)
	msg.SetBody("text/plain", notify.Body)

	d := gomail.NewDialer("smtp.mail.ru", 465, from, os.Getenv("EMAIL_PASSWORD"))

	if err := d.DialAndSend(msg); err != nil {
		elk.Log.SendMsg(
			elk.LogMessage{
				Level:   'E',
				Message: "Error send notify",
				Fields: map[string]interface{}{
					"method": "sendNotify",
					"action": "DialAndSend",
					"error":  err,
				},
			})
		return err
	}
	elk.Log.SendMsg(
		elk.LogMessage{
			Level:   'I',
			Message: "Send notify",
			Fields: map[string]interface{}{
				"method":  "sendNotify",
				"action":  "DialAndSend",
				"email":   notify.Email,
				"subject": notify.Subject,
				"body":    notify.Body,
			},
		})

	return nil
}

func (s *ServiceManager) workerSendNotify(ctx context.Context) {
	for {
		select {
		case notify := <-s.notifyChan:
			if err := s.sendNotify(notify); err != nil {
				elk.Log.SendMsg(
					elk.LogMessage{
						Level:   'E',
						Message: "Error send notify",
						Fields: map[string]interface{}{
							"method": "workerSendNotify",
							"action": "sendNotify",
							"error":  err,
						},
					})
			}
		case <-ctx.Done():
			return
		}
	}
}

func (s *ServiceManager) InputNotify(notify models.Notification) {
	s.notifyChan <- notify
}

func (s *ServiceManager) Close() {
	close(s.notifyChan)
}

func (s *ServiceManager) createWorkers(ctx context.Context, countWorker int) {
	for i := 0; i < countWorker; i++ {
		go s.workerSendNotify(ctx)
	}
}

func (s *ServiceManager) Start(ctx context.Context, countWorker int) {
	s.createWorkers(ctx, countWorker)
}
