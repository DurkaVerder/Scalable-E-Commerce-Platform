package consumer

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/notification-service/internal/kafka"
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/notification-service/internal/models"
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/notification-service/internal/service"
	elk "github.com/DurkaVerder/elk-send-logs/elk"

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
			elk.Log.SendMsg(
				elk.LogMessage{
					Level:   'I',
					Message: "Consumer created successfully",
					Fields: map[string]interface{}{
						"method": "NewConsumerManager",
						"action": "NewConsumer",
					},
				})

			return &ConsumerManager{
				consumer: consumer,
				config:   *config,
				service:  service,
			}
		}
		elk.Log.SendMsg(elk.LogMessage{
			Level:   'E',
			Message: "Failed to create consumer",
			Fields: map[string]interface{}{
				"method": "NewConsumerManager",
				"action": "NewConsumer",
				"error":  err,
			},
		})
		time.Sleep(time.Second * 2)
	}

	elk.Log.SendMsg(elk.LogMessage{
		Level:   'E',
		Message: "Failed to create consumer after retries",
		Fields: map[string]interface{}{
			"method": "NewConsumerManager",
			"action": "NewConsumer",
		},
	})

	log.Fatalln("Failed to create consumer")
	return nil
}

func (c *ConsumerManager) Subscribe(topic string) error {
	var err error
	c.consumePartition, err = c.consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		elk.Log.SendMsg(elk.LogMessage{
			Level:   'E',
			Message: "Failed to consume partition",
			Fields: map[string]interface{}{
				"method": "Subscribe",
				"action": "ConsumePartition",
				"error":  err,
			},
		})
		return err
	}
	return nil
}

func (c *ConsumerManager) Start(ctx context.Context) {
	for {
		select {
		case msg := <-c.consumePartition.Messages():
			elk.Log.SendMsg(elk.LogMessage{
				Level:   'I',
				Message: "Message received",
				Fields: map[string]interface{}{
					"method":  "Start",
					"action":  "Messages",
					"message": string(msg.Value),
				},
			})

			notify := models.Notification{}
			if err := json.Unmarshal(msg.Value, &notify); err != nil {
				elk.Log.SendMsg(elk.LogMessage{
					Level:   'E',
					Message: "Error unmarshalling message",
					Fields: map[string]interface{}{
						"method": "Start",
						"action": "Unmarshal",
						"error":  err,
					},
				})
				continue
			}
			c.service.InputNotify(notify)

		case err := <-c.consumePartition.Errors():
			elk.Log.SendMsg(elk.LogMessage{
				Level:   'E',
				Message: "Error consuming partition",
				Fields: map[string]interface{}{
					"method": "Start",
					"action": "Errors",
					"error":  err,
				},
			})

		case <-ctx.Done():
			elk.Log.SendMsg(elk.LogMessage{
				Level:   'I',
				Message: "Context done",
				Fields: map[string]interface{}{
					"method": "Start",
					"action": "Done",
				},
			})
			return
		}
	}
}

func (c *ConsumerManager) Close() {
	if err := c.consumePartition.Close(); err != nil {
		elk.Log.SendMsg(elk.LogMessage{
			Level:   'E',
			Message: "Failed to close consume partition",
			Fields: map[string]interface{}{
				"method": "Close",
				"action": "Close",
				"error":  err,
			},
		})
		panic(err)
	}
	if err := c.consumer.Close(); err != nil {
		elk.Log.SendMsg(elk.LogMessage{
			Level:   'E',
			Message: "Failed to close consumer",
			Fields: map[string]interface{}{
				"method": "Close",
				"action": "Close",
				"error":  err,
			},
		})
		panic(err)
	}

	elk.Log.SendMsg(elk.LogMessage{
		Level:   'I',
		Message: "Consumer closed successfully",
		Fields: map[string]interface{}{
			"method": "Close",
			"action": "Close",
		},
	})
}
