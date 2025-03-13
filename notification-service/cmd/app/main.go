package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/notification-service/internal/kafka/consumer"
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/notification-service/internal/service"
	elk "github.com/DurkaVerder/Scalable-E-Commerce-Platform/notification-service/pkg/logs"
)

func main() {
	sizeChan := getIntEnv("SIZE_CHAN")
	service := service.NewServiceManager(sizeChan)

	consumer := consumer.NewConsumerManager([]string{os.Getenv("KAFKA_BROKER")}, service)
	consumer.Subscribe(os.Getenv("KAFKA_TOPIC"))

	ctx := context.Background()
	countWorker := getIntEnv("COUNT_WORKER")

	go service.Start(ctx, countWorker)
	go consumer.Start(ctx)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-signalChan

	elk.Log.Info("Signal received", map[string]interface{}{
		"method": "main",
		"action": "signal received",
		"signal": sig,
	})

	consumer.Close()
	service.Close()
}

func getIntEnv(key string) int {
	val, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		elk.Log.Error("Failed to convert env to int", map[string]interface{}{
			"method": "main",
			"action": "get int env",
			"key":    key,
		})
		log.Fatalf("Failed to convert %s to int", key)
	}
	return val
}
