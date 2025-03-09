package server

import (
	"cart-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

type Server struct {
	handlers handlers.Handlers
	r        *gin.Engine
}

func NewServer(handlers handlers.Handlers, engine *gin.Engine) *Server {
	return &Server{
		handlers: handlers,
		r:        engine,
	}
}

func (s *Server) configureRouter() {
	cart := s.r.Group("/cart")
	{
		cart.POST("/add", s.handlers.HandlerAddProduct)
		cart.GET("/get", s.handlers.HandlerGetCart)
		cart.DELETE("/delete/:product_id", s.handlers.HandlerDeleteProduct)
		cart.PUT("/update/:product_id", s.handlers.HandlerUpdateProduct)
	}
}

func (s *Server) Start() {
	s.configureRouter()
	s.r.Run(":8081")
}
