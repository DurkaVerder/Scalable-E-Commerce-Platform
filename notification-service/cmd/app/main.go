package main

import (
	"context"
	"os"
	"strconv"

	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/notification-service/internal/kafka/consumer"
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/notification-service/internal/service"
	elk "github.com/DurkaVerder/elk-send-logs/elk"
)

func main() {
	elk.InitLogger(5, "notification-service", os.Getenv("ELK_URL"))
	ctx := context.Background()
	elk.Log.Start(ctx, 3)

	sizeChan := getIntEnv("SIZE_CHAN")
	service := service.NewServiceManager(sizeChan)

	consumer := consumer.NewConsumerManager([]string{os.Getenv("KAFKA_BROKER")}, service)
	consumer.Subscribe(os.Getenv("KAFKA_TOPIC"))

	countWorker := getIntEnv("COUNT_WORKER")

	go service.Start(ctx, countWorker)
	go consumer.Start(ctx)

	<-ctx.Done()

	consumer.Close()
	service.Close()
}

func getIntEnv(key string) int {
	val, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		elk.Log.SendMsg(
			elk.LogMessage{
				Level:   'E',
				Message: "Failed to convert to int",
				Fields: map[string]interface{}{
					"method": "getIntEnv",
					"action": "Atoi",
					"error":  err,
				},
			})
		panic(err)
	}
	return val
}
