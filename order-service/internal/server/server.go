package server

import (
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/order-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

type Server struct {
	handlers handlers.Handler
	r        *gin.Engine
}

func NewServer(handlers handlers.Handler) *Server {
	return &Server{
		handlers: handlers,
	}
}

func (s *Server) initRouters() {
	s.r = gin.Default()

	order := s.r.Group("/order")
	{
		order.GET("/:orderID", s.handlers.HandlerGetOrder)
		order.GET("/all/:userID", s.handlers.HandlerGetOrders)
		order.POST("/:userID", s.handlers.HandlerCreateOrder)
		order.PUT("/:orderID", s.handlers.HandlerUpdateOrder)
	}

}

func (s *Server) Run(port string) {
	s.initRouters()

	if err := s.r.Run(port); err != nil {
		panic(err)
	}
}
