package controller

import (
	"portfolio-web-be/database"
	model "portfolio-web-be/models"

	"github.com/gofiber/fiber/v2"
)

func CheckUser(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)

	var user model.User
	if err := database.DB.First(&user, userId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.JSON(user)
}