package main

import (
	"os"

	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/catalog-service/internal/handlers"
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/catalog-service/internal/repository/postgres"
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/catalog-service/internal/server"
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/catalog-service/internal/service"
	elk "github.com/DurkaVerder/Scalable-E-Commerce-Platform/catalog-service/pkg/logs"
)

func main() {
	db, err := postgres.ConnectDB(os.Getenv("DB_URL"))
	if err != nil {
		elk.Log.Error("Error connecting to DB", map[string]interface{}{
			"method": "main",
			"action": "ConnectDB",
			"error":  err.Error(),
		})
		panic(err)
	}

	postgres := postgres.NewPostgres(db)

	service := service.NewServiceManager(postgres)

	handlers := handlers.NewHandlersManager(service)

	server := server.NewServer(handlers)

	server.Start(os.Getenv("PORT"))

	elk.Log.Info("Server started", map[string]interface{}{
		"method": "main",
		"action": "Start",
	})

}
