package main

import (
	"context"
	"os"

	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/order-service/internal/handlers"
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/order-service/internal/repository/postgres"
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/order-service/internal/server"
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/order-service/internal/service"
	"github.com/DurkaVerder/elk-send-logs/elk"
)

func main() {
	elk.InitLogger(5, "order-service", os.Getenv("ELK_URL"))
	ctx := context.Background()
	elk.Log.Start(ctx, 3)

	db := postgres.ConnectDB(os.Getenv("DB_URL"))

	postgres := postgres.NewPostgres(db)
	defer db.Close()

	service := service.NewServiceManager(postgres)

	handlers := handlers.NewHandlerManager(service)

	server := server.NewServer(handlers)

	go server.Run(os.Getenv("PORT"))

	<-ctx.Done()

	elk.Log.Close()
}
