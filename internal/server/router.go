package server

import (
	"github.com/gofiber/fiber/v3"
	"github.com/odysseymorphey/SimpleAuth/internal/repository"
	"github.com/odysseymorphey/SimpleAuth/internal/server/handlers"
	"github.com/sirupsen/logrus"
)

type Router struct {
	baseHandler *handlers.BaseHandler
}

func NewRouter(repo repository.Repository, logger *logrus.Logger) *Router {
	return &Router{
		baseHandler: handlers.NewBaseHandler(repo, logger),
	}
}

func (r *Router) RegisterRoutes(srv *fiber.App) {
	api := srv.Group("/api")

	{
		api.Post("/token", r.baseHandler.GenerateToken)
		api.Post("/refresh", r.baseHandler.RefreshToken)
	}
}
