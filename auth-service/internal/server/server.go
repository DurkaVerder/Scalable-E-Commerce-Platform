package server

import (
	"auth-service/internal/handlers"

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

func (s *Server) Run() {
	s.initRoutes()

	if err := s.r.Run(":8080"); err != nil {
		panic(err)
	}
}
