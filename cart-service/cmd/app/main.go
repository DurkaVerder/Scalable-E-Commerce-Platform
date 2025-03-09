package main

import (
	"os"

	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/cart-service/internal/handlers"
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/cart-service/internal/repository/postgres"
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/cart-service/internal/server"
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/cart-service/internal/service"
	elk "github.com/DurkaVerder/Scalable-E-Commerce-Platform/cart-service/pkg/logs"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := postgres.ConnectDB(os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}

	repo := postgres.NewPostgres(db)

	service := service.NewServiceManager(repo)

	handlers := handlers.NewHandlersManager(service)

	r := gin.Default()

	server := server.NewServer(handlers, r)

	server.Start(os.Getenv("PORT"))

	elk.Log.Info("Server started at port", map[string]interface{}{
		"method": "Start",
		"action": "starting server",
	})
}
