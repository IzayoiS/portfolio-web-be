package routes

import (
	controller "portfolio-web-be/controllers"

	"github.com/gofiber/fiber/v2"
)

func LoginRoutes(app *fiber.App) {
	app.Post("/login", controller.Login)
}