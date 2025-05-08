package routes

import (
	controller "portfolio-web-be/controllers"
	"portfolio-web-be/middleware"

	"github.com/gofiber/fiber/v2"
)

func TechRoutes(router fiber.Router) {
	stack := router.Group("/tech", middleware.Protected)
	stack.Post("/", controller.CreateTechStack)
	stack.Get("/", controller.GetTechStacks)
}