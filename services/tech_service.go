package service

import (
	"portfolio-web-be/database"
	model "portfolio-web-be/models"

	"github.com/gofiber/fiber/v2"
)

func CreateTechStack(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)

	var body struct {
		Name string `json:"name"`
	}
	if err := c.BodyParser(&body); err != nil || body.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Name is required"})
	}

	tech := model.TechStack{
		UserID: userId,
		Name:   body.Name,
	}
	if err := database.DB.Create(&tech).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create tech stack"})
	}
	return c.JSON(tech)
}

func GetTechStacks(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)
	var stacks []model.TechStack
	if err := database.DB.Preload("User").Where("user_id = ?", userId).Find(&stacks).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Fetch tech stacks failed"})
	}
	return c.JSON(stacks)
}