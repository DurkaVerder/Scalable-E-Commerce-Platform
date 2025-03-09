package main

import (
	"cart-service/internal/handlers"
	elk "cart-service/pkg/logs"
	"cart-service/internal/repository/postgres"
	"cart-service/internal/server"
	"cart-service/internal/service"
	"os"

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
