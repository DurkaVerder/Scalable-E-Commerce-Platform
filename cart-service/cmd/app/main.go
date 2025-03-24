package main

import (
	"context"
	"os"

	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/cart-service/internal/handlers"
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/cart-service/internal/repository/postgres"
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/cart-service/internal/server"
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/cart-service/internal/service"
	elk "github.com/DurkaVerder/elk-send-logs/elk"
	"github.com/gin-gonic/gin"
)

func main() {
	elk.InitLogger(10, "cart-service", "logstash:5000")

	ctx := context.Background()

	go elk.Log.Start(ctx, 3)

	db, err := postgres.ConnectDB(os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}

	repo := postgres.NewPostgres(db)

	service := service.NewServiceManager(repo)

	handlers := handlers.NewHandlersManager(service)

	r := gin.Default()

	server := server.NewServer(handlers, r)

	go server.Start(os.Getenv("PORT"))

	elk.Log.SendMsg(elk.LogMessage{
		Level:   'I',
		Message: "Server started",
		Fields: map[string]interface{}{
			"method": "main",
			"action": "start",
			"port":   os.Getenv("PORT"),
		},
	})

	<-ctx.Done()

	elk.Log.SendMsg(elk.LogMessage{
		Level:   'I',
		Message: "Server stopped",
		Fields: map[string]interface{}{
			"method": "main",
			"action": "stop",
		},
	})

	elk.Log.Close()
}
