package main

import (
	"os"

	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/payment-service/internal/handlers"
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/payment-service/internal/repository/postgres"
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/payment-service/internal/server"
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/payment-service/internal/service"
)

func main() {
	db := postgres.ConnectDB(os.Getenv("DB_URL"))

	postgres := postgres.NewPostgres(db)

	defer postgres.Close()

	service := service.NewPaymentService(postgres)

	handlers := handlers.NewHandlerManager(service)

	server := server.NewServer(handlers)

	server.Run(os.Getenv("PORT"))

}
