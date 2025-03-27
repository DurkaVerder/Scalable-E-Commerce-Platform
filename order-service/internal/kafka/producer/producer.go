// This package contains the Producer interface and ProducerManager struct, which are used to send messages to Kafka topics.
package producer

import (
	"encoding/json"
	"time"

	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/order-service/internal/kafka"
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/order-service/internal/models"
	"github.com/DurkaVerder/elk-send-logs/elk"
	"github.com/IBM/sarama"
)

// Producer is a wrapper around the sarama.SyncProducer to provide a more
type Producer interface {
	SendMessage(topic string, order models.Notification, maxRetry int) error
}

// ProducerManager is a wrapper around the sarama.SyncProducer to provide a more
type ProducerManager struct {
	producer sarama.SyncProducer
	config   *sarama.Config
}

// NewProducerManager creates a new ProducerManager using the given broker addresses.
func NewProducerManager(brokers string) *ProducerManager {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = kafka.MaxRetries

	for i := 0; i < kafka.MaxRetries; i++ {

		producer, err := sarama.NewSyncProducer([]string{brokers}, config)
		if err == nil {
			elk.Log.SendMsg(
				elk.LogMessage{
					Level:   'I',
					Message: "Producer created",
					Fields: map[string]interface{}{
						"method":  "NewProducerManager",
						"action":  "Producer created",
						"brokers": brokers,
					},
				})
			return &ProducerManager{producer, config}
		}
		elk.Log.SendMsg(
			elk.LogMessage{
				Level:   'E',
				Message: "Error creating producer",
				Fields: map[string]interface{}{
					"method":  "NewProducerManager",
					"error":   err.Error(),
					"brokers": brokers,
				},
			})
		time.Sleep(5 * time.Second)
	}
	elk.Log.SendMsg(
		elk.LogMessage{
			Level:   'F',
			Message: "Error creating producer after max retries",
			Fields: map[string]interface{}{
				"method":  "NewProducerManager",
				"brokers": brokers,
			},
		})
	return nil
}

// SendMessageForAddOrder sends a message to the Kafka topic for adding an order.
func (p *ProducerManager) SendMessage(topic string, order models.Notification, maxRetry int) error {
	data, err := json.Marshal(order)
	if err != nil {
		elk.Log.SendMsg(
			elk.LogMessage{
				Level:   'E',
				Message: "Error marshaling order",
				Fields: map[string]interface{}{
					"method": "SendMessage",
					"error":  err.Error(),
				},
			})
		return nil
	}

	for i := 0; i < maxRetry; i++ {
		if err = p.sendMessage(topic, data); err == nil {
			return nil
		}
		elk.Log.SendMsg(
			elk.LogMessage{
				Level:   'E',
				Message: "Error sending message to Kafka",
				Fields: map[string]interface{}{
					"method":  "SendMessage",
					"retry":   i + 1,
					"error":   err.Error(),
					"topic":   topic,
					"payload": string(data),
				},
			})
	}

	return err
}

// sendMessage sends a message to the Kafka topic.
func (p *ProducerManager) sendMessage(topic string, data []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(data),
	}

	partition, offset, err := p.producer.SendMessage(msg)
	if err != nil {
		elk.Log.SendMsg(
			elk.LogMessage{
				Level:   'E',
				Message: "Error sending message to Kafka",
				Fields: map[string]interface{}{
					"method":  "sendMessage",
					"error":   err.Error(),
					"topic":   topic,
					"payload": string(data),
				},
			})
		return err
	}

	elk.Log.SendMsg(
		elk.LogMessage{
			Level:   'I',
			Message: "Message sent to Kafka",
			Fields: map[string]interface{}{
				"method":    "sendMessage",
				"topic":     topic,
				"partition": partition,
				"offset":    offset,
			},
		})
	return nil
}
