package routes

import (
	controller "portfolio-web-be/controllers"
	"portfolio-web-be/middleware"

	"github.com/gofiber/fiber/v2"
)

func TechRoutes(router fiber.Router) {
	stack := router.Group("/tech")
	stack.Post("/",middleware.Protected, controller.CreateTechStack)
	stack.Get("/", controller.GetTechStacks)
	stack.Patch("/:id",middleware.Protected, controller.UpdateTechStacks)
	stack.Delete("/:id",middleware.Protected, controller.DeleteTechStacks)
}