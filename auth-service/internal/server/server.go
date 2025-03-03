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

func (s *Server) init() {
	s.r = gin.Default()

	s.r.POST("/login", s.handlers.HandlerLogin)
	s.r.POST("/register", s.handlers.HandlerRegister)
}

func (s *Server) Run() {
	s.init()

	if err := s.r.Run(); err != nil {
		panic(err)
	}
}
