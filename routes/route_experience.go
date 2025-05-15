package routes

import (
	controller "portfolio-web-be/controllers"
	"portfolio-web-be/middleware"

	"github.com/gofiber/fiber/v2"
)

func ExperienceRoutes(router fiber.Router) {
	exp := router.Group("/experience", middleware.Protected)
	exp.Post("/", controller.CreateExperience)
	exp.Get("/", controller.GetExperience)
	exp.Patch("/:id", controller.UpdateExperience)
	exp.Delete("/:id", controller.DeleteExperience)
}