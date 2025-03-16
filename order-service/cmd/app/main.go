package main

import (
	"os"

	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/order-service/internal/handlers"
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/order-service/internal/repository/postgres"
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/order-service/internal/server"
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/order-service/internal/service"
)

func main() {
	db := postgres.ConnectDB(os.Getenv("URL_DB"))

	postgres := postgres.NewPostgres(db)
	defer db.Close()

	service := service.NewServiceManager(postgres)

	handlers := handlers.NewHandlerManager(service)

	server := server.NewServer(handlers)


	server.Run(os.Getenv("PORT"))
}
