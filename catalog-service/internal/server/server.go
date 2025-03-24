package server

import (
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/catalog-service/internal/handlers"
	"github.com/DurkaVerder/elk-send-logs/elk"
	"github.com/gin-gonic/gin"
)

type Server struct {
	handlers handlers.Handlers
	r        *gin.Engine
}

func NewServer(handlers handlers.Handlers) *Server {
	return &Server{handlers: handlers}
}

func (s *Server) initRoutes() {
	s.r = gin.Default()

	catalog := s.r.Group("/catalog")
	{
		catalog.GET("/products", s.handlers.HandlerAllProducts)
		catalog.GET("/products/:id", s.handlers.HandlerProductByID)
	}
}

func (s *Server) Start(port string) {
	s.initRoutes()

	if err := s.r.Run(port); err != nil {
		elk.Log.SendMsg(
			elk.LogMessage{
				Level:   'E',
				Message: "Failed to start server",
				Fields: map[string]interface{}{
					"method": "Start",
					"error":  err,
				},
			})

		panic(err)
	}
}
