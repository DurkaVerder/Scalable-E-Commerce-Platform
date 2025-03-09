package server

import (
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/cart-service/internal/handlers"
	elk "github.com/DurkaVerder/Scalable-E-Commerce-Platform/cart-service/pkg/logs"

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
	if err := s.r.Run(port); err != nil {
		elk.Log.Error("Server started at port "+port, map[string]interface{}{
			"method": "Start",
			"action": "starting server",
			"error":  err,
			"port":   port,
		})
		panic(err)
	}

}
