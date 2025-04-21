package handlers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"github.com/odysseymorphey/SimpleAuth/internal/models"
	"github.com/odysseymorphey/SimpleAuth/internal/repository"
	"github.com/odysseymorphey/SimpleAuth/internal/services"
	"github.com/sirupsen/logrus"
)

type BaseHandler struct {
	repo   repository.Repository
	logger *logrus.Logger
}

func NewBaseHandler(r repository.Repository, logger *logrus.Logger) *BaseHandler {
	return &BaseHandler{
		repo:   r,
		logger: logger,
	}
}

func (h *BaseHandler) GenerateToken(c fiber.Ctx) error {
	guid := c.Query("guid")
	if guid == "" {
		return c.JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"status":  "Bad Request",
			"message": "GUID is empty",
		})
	}

	userInfo := &models.UserInfo{
		GUID:   guid,
		UserIP: c.IP(),
	}

	tokenPair, err := services.GeneratePair(h.repo, userInfo)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"status":  "Internal Server Error",
			"message": "Cannot generate token pair",
		})
	}

	return c.JSON(fiber.Map{
		"code":     fiber.StatusOK,
		"status":   "OK",
		"authData": tokenPair,
	})
}

func (h *BaseHandler) RefreshToken(c fiber.Ctx) error {
	guid := c.Query("guid")
	if guid == "" {
		return c.JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"status":  "Bad Request",
			"message": "GUID is empty",
		})
	}

	tokenPair := &models.Pair{}
	if err := json.Unmarshal(c.Body(), tokenPair); err != nil {
		return c.JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"status":  "Bad Request",
			"message": "invalid request body data",
		})
	}

	uInfo := &models.UserInfo{
		GUID:   guid,
		UserIP: c.IP(),
	}

	newPair, err := services.RefreshAccessToken(h.repo, uInfo, tokenPair)
	if err != nil {
		h.logger.Error(err)

		return c.JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"status":  "Internal Server Error",
			"message": "cannot refresh tokens",
		})
	}

	return c.JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "OK",
		"tokens": newPair,
	})
}
