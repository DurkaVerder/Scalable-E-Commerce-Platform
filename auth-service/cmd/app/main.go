package main

import (
	"context"
	"os"

	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/auth-service/internal/handlers"
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/auth-service/internal/repository/postgres"
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/auth-service/internal/server"
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/auth-service/internal/service"
	"github.com/DurkaVerder/elk-send-logs/elk"
)

func main() {
	elk.InitLogger(5, "auth-service", os.Getenv("ELK_URL"))

	ctx := context.Background()

	elk.Log.Start(ctx, 5)

	db := postgres.ConnectToDB()

	postgres := postgres.NewPostgres(db)

	service := service.NewServiceManager(postgres)

	handlers := handlers.NewHandlersManager(service)

	server := server.NewServer(handlers)

	go server.Start(os.Getenv("PORT"))

	elk.Log.SendMsg(
		elk.LogMessage{
			Level:   'I',
			Message: "Server started",
			Fields: map[string]interface{}{
				"method": "main",
				"action": "start",
				"port":   os.Getenv("PORT"),
			},
		})

	<-ctx.Done()

	elk.Log.Close()
	postgres.Close()
}
