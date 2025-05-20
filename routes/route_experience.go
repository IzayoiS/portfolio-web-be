package routes

import (
	controller "portfolio-web-be/controllers"
	"portfolio-web-be/middleware"

	"github.com/gofiber/fiber/v2"
)

func ExperienceRoutes(router fiber.Router) {
	exp := router.Group("/experience")
	exp.Post("/",middleware.Protected, controller.CreateExperience)
	exp.Get("/", controller.GetExperience)
	exp.Patch("/:id",middleware.Protected, controller.UpdateExperience)
	exp.Delete("/:id",middleware.Protected, controller.DeleteExperience)
}