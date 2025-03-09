package server

import (
	"cart-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

// Server represents the server
type Server struct {
	handlers handlers.Handlers
	r        *gin.Engine
}

// NewServer creates a new server
func NewServer(handlers handlers.Handlers, engine *gin.Engine) *Server {
	return &Server{
		handlers: handlers,
		r:        engine,
	}
}

// configureRouter configures the routes
func (s *Server) configureRouter() {
	cart := s.r.Group("/cart")
	{
		cart.POST("/add", s.handlers.HandlerAddProduct)
		cart.GET("/get", s.handlers.HandlerGetCart)
		cart.DELETE("/delete/:product_id", s.handlers.HandlerDeleteProduct)
		cart.PUT("/update/:product_id", s.handlers.HandlerUpdateProduct)
	}
}

// Start starts the server
func (s *Server) Start(port string) {
	s.configureRouter()
	s.r.Run(port)
}
