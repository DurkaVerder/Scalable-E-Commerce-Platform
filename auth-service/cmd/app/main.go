package main

import (
	"os"

	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/auth-service/internal/handlers"
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/auth-service/internal/repository/postgres"
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/auth-service/internal/server"
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/auth-service/internal/service"
)

func main() {
	db, err := postgres.ConnectToDB()
	if err != nil {
		panic(err)
	}

	postgres := postgres.NewPostgres(db)

	service := service.NewServiceManager(postgres)

	handlers := handlers.NewHandlersManager(service)

	server := server.NewServer(handlers)

	server.Start(os.Getenv("PORT"))
}
