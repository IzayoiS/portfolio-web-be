package routes

import (
	controller "portfolio-web-be/controllers"
	"portfolio-web-be/middleware"

	"github.com/gofiber/fiber/v2"
)

func ProfileRoute(router fiber.Router) {
	profile := router.Group("/profile", middleware.Protected)
	profile.Post("/",controller.CreateProfile)
	profile.Get("/:userId",controller.GetProfile)
	profile.Patch("/:userId",controller.UpdateProfile)
}
