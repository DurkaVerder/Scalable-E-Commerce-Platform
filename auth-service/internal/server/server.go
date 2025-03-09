package server

import (
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/auth-service/internal/handlers"

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
	auth := s.r.Group("/auth")
	auth.POST("/login", s.handlers.HandlerLogin)
	auth.POST("/register", s.handlers.HandlerRegister)
	auth.GET("/validate", s.handlers.HandlerValidateToken)
}

func (s *Server) Start(port string) {
	s.initRoutes()

	if err := s.r.Run(port); err != nil {
		panic(err)
	}
}
