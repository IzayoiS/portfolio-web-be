package routes

import (
	controller "portfolio-web-be/controllers"
	"portfolio-web-be/middleware"

	"github.com/gofiber/fiber/v2"
)

func ProjectRoute(router fiber.Router) {
	project := router.Group("/project")
	project.Post("/",middleware.Protected, controller.CreateProject)
	project.Get("/", controller.GetProject)
	project.Patch("/:id",middleware.Protected, controller.UpdateProject)
	project.Delete("/:id",middleware.Protected, controller.DeleteProject)
}