package controller

import (
	service "portfolio-web-be/services"

	"github.com/gofiber/fiber/v2"
)

func CreateProject(c *fiber.Ctx) error {
	return service.CreateProject(c)
}

func GetProject(c *fiber.Ctx) error {
	return service.GetProject(c)
}

func UpdateProject(c *fiber.Ctx) error {
	return service.UpdateProject(c)
}

func DeleteProject(c *fiber.Ctx) error {
	return service.DeleteProject(c)
}