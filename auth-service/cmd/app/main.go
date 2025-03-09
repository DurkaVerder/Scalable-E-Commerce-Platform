package main

import (
	"auth-service/internal/handlers"
	"auth-service/internal/repository/postgres"
	"auth-service/internal/server"
	"auth-service/internal/service"
	"os"
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
