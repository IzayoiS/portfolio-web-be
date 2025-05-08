package controller

import (
	service "portfolio-web-be/services"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var body struct {
		Username     string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	
	user, err := service.RegisterUser(body.Username, body.Email, body.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "User registered successfully",
		"user": fiber.Map{
			"id":    user.ID,
			"username":  user.Username,
			"email": user.Email,
		},
	})
}
