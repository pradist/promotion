package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pradist/promotion/handlers"
	"github.com/pradist/promotion/repositories"
	"github.com/pradist/promotion/services"
)

func New() *fiber.App {
	r := repositories.NewPromotionRepository()
	s := services.NewPromotionService(r)
	h := handlers.NewPromotionHandler(s)

	app := fiber.New()
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("OK")
	})

	app.Get("/api/v1/promotion", h.CalculateDiscount)
	return app
}
