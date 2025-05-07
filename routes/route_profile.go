package routes

import (
	controller "portfolio-web-be/controllers"

	"github.com/gofiber/fiber/v2"
)

func ProfileRoute(app *fiber.App) {
	app.Post("/profile",controller.CreateProfile)
	app.Get("/profile/:userId",controller.GetProfile)
	app.Patch("/profile/:userId",controller.UpdateProfile)
}