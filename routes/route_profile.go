package routes

import (
	controller "portfolio-web-be/controllers"
	"portfolio-web-be/middleware"

	"github.com/gofiber/fiber/v2"
)

func ProfileRoute(router fiber.Router) {
	profile := router.Group("/profile")
	profile.Post("/",middleware.Protected, controller.CreateProfile)
	profile.Get("/:userId",controller.GetProfile)
	profile.Patch("/:userId",middleware.Protected, controller.UpdateProfile)
}
