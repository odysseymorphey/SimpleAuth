package server

import (
	"github.com/gofiber/fiber/v3"
	"github.com/odysseymorphey/SimpleAuth/internal/repository"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	server *fiber.App
	router *Router
	repo   repository.Repository
	logger *logrus.Logger
}

func NewServer(repo repository.Repository) *Server {
	srv := fiber.New()

	logger := logrus.New()

	r := NewRouter(repo, logger)

	r.RegisterRoutes(srv)

	return &Server{
		server: srv,
		router: r,
		repo:   repo,
		logger: logger,
	}
}

func (s *Server) Start() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)

	go func() {
		<-sig
		s.Destroy()
		s.logger.Info("Server stopped")
		os.Exit(0)
	}()

	s.logger.Fatal(s.server.Listen(":8080"))

	log.Println("Server started on port 8080")
}

func (s *Server) Destroy() {
	s.repo.Close()
}
