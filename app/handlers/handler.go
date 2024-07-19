package handlers

import (
	"try-go-vercels/app/services"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(app *fiber.App) {
	app.Get("/", h.Home)
}

func (h *Handler) Home(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
