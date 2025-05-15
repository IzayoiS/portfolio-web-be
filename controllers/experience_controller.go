package controller

import (
	service "portfolio-web-be/services"

	"github.com/gofiber/fiber/v2"
)

func CreateExperience(c *fiber.Ctx) error {
	return service.CreateExperience(c)
}

func GetExperience(c *fiber.Ctx) error {
	return service.GetExperience(c)
}

func UpdateExperience(c *fiber.Ctx) error {
	return service.UpdateExperience(c)
}

func DeleteExperience(c *fiber.Ctx) error {
	return service.DeleteExperience(c)
}