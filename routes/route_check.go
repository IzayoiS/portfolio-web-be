package routes

import (
	controller "portfolio-web-be/controllers"
	"portfolio-web-be/middleware"

	"github.com/gofiber/fiber/v2"
)

func CheckUserRoutes(app *fiber.App) {
	app.Get("/check", middleware.Protected, controller.CheckUser)
}