package server

import (
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/payment-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

type Server struct {
	handlers handlers.Handler
	r        *gin.Engine
}

func NewServer(handlers handlers.Handler) *Server {
	return &Server{handlers: handlers}
}

func (s *Server) initRoutes() {
	s.r = gin.Default()

	payment := s.r.Group("/payment")
	{
		payment.POST("/pay", s.handlers.HandlerCreatePayment)
	}
}

func (s *Server) Run(port string) {
	s.initRoutes()

	if err := s.r.Run(port); err != nil {
		panic(err)
	}
}
