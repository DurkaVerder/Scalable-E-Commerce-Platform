package consumer

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/notification-service/internal/kafka"
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/notification-service/internal/models"
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/notification-service/internal/service"
	elk "github.com/DurkaVerder/Scalable-E-Commerce-Platform/notification-service/pkg/logs"

	"github.com/IBM/sarama"
)

type Consumer interface {
	Subscribe(topic string) error
	Start()
}

type ConsumerManager struct {
	consumer         sarama.Consumer
	config           sarama.Config
	consumePartition sarama.PartitionConsumer
	service          service.Service
}

func NewConsumerManager(brokers []string, service service.Service) *ConsumerManager {
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	for i := 0; i < kafka.MaxRetry; i++ {
		consumer, err := sarama.NewConsumer(brokers, config)
		if err == nil {
			elk.Log.Info("Create consumer", map[string]interface{}{
				"method": "NewConsumerManager",
				"action": "NewConsumer",
			})

			return &ConsumerManager{
				consumer: consumer,
				config:   *config,
				service:  service,
			}
		}
		elk.Log.Error("Failed to create consumer", map[string]interface{}{
			"method": "NewConsumerManager",
			"action": "NewConsumer",
			"error":  err,
		})
		time.Sleep(time.Second * 2)
	}

	elk.Log.Error("Failed to create consumer", map[string]interface{}{
		"method": "NewConsumerManager",
		"action": "NewConsumer",
	})

	log.Fatalln("Failed to create consumer")
	return nil
}

func (c *ConsumerManager) Subscribe(topic string) error {
	var err error
	c.consumePartition, err = c.consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		elk.Log.Error("Failed to consume partition", map[string]interface{}{
			"method": "Subscribe",
			"action": "ConsumePartition",
			"error":  err,
		})
		return err
	}
	return nil
}

func (c *ConsumerManager) Start(ctx context.Context) {
	for {
		select {
		case msg := <-c.consumePartition.Messages():
			elk.Log.Info("Receive message", map[string]interface{}{
				"method":  "Start",
				"action":  "Messages",
				"message": string(msg.Value),
			})

			notify := models.Notification{}
			if err := json.Unmarshal(msg.Value, &notify); err != nil {
				elk.Log.Error("Error unmarshal message", map[string]interface{}{
					"method": "Start",
					"action": "Unmarshal",
					"error":  err,
				})
				continue
			}
			c.service.InputNotify(notify)

		case err := <-c.consumePartition.Errors():
			elk.Log.Error("Error consume partition", map[string]interface{}{
				"method": "Start",
				"action": "Errors",
				"error":  err,
			})

		case <-ctx.Done():
			elk.Log.Info("Context done", map[string]interface{}{
				"method": "Start",
				"action": "Done",
			})
			return
		}
	}
}

func (c *ConsumerManager) Close() {
	if err := c.consumePartition.Close(); err != nil {
		elk.Log.Error("Failed to close consume partition", map[string]interface{}{
			"method": "Close",
			"action": "Close",
			"error":  err,
		})
		panic(err)
	}
	if err := c.consumer.Close(); err != nil {
		elk.Log.Error("Failed to close consumer", map[string]interface{}{
			"method": "Close",
			"action": "Close",
			"error":  err,
		})
		panic(err)
	}

	elk.Log.Info("Close consumer", map[string]interface{}{
		"method": "Close",
		"action": "Close",
	})
}
