package controller

import (
	service "portfolio-web-be/services"

	"github.com/gofiber/fiber/v2"
)

func CreateTechStack(c *fiber.Ctx) error {
	return service.CreateTechStack(c)
}

func GetTechStacks(c *fiber.Ctx) error {
	return service.GetTechStacks(c)
}

func UpdateTechStacks(c *fiber.Ctx) error {
	return service.UpdateTech(c)
}

func DeleteTechStacks(c *fiber.Ctx) error {
	return service.DeleteTech(c)
}