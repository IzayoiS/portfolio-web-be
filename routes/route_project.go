package routes

import (
	controller "portfolio-web-be/controllers"
	"portfolio-web-be/middleware"

	"github.com/gofiber/fiber/v2"
)

func ProjectRoute(router fiber.Router) {
	project := router.Group("/project", middleware.Protected)
	project.Post("/", controller.CreateProject)
	project.Get("/", controller.GetProject)
	project.Patch("/:id", controller.UpdateProject)
	project.Delete("/:id", controller.DeleteProject)
}