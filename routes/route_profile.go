package routes

import (
	controller "portfolio-web-be/controllers"

	"github.com/gofiber/fiber/v2"
)

func ProfileRoute(router fiber.Router) {
	profile := router.Group("/profile")
	profile.Post("/",controller.CreateProfile)
	profile.Get("/:userId",controller.GetProfile)
	profile.Patch("/:userId",controller.UpdateProfile)
}
