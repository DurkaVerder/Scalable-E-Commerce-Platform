package main

import (
	"auth-service/internal/handlers"
	"auth-service/internal/repository/postgres"
	"auth-service/internal/server"
	"auth-service/internal/service"
)

func main() {
	postgres := postgres.NewPostgres(nil)

	service := service.NewServiceManager(postgres)

	handlers := handlers.NewHandlersManager(service)

	server := server.NewServer(handlers)

	server.Run()
}
