package routes

import (
	controller "portfolio-web-be/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoute(app *fiber.App) {
	app.Post("/register", controller.Register)
}