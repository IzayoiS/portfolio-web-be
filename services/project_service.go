package service

import (
	"encoding/json"
	"portfolio-web-be/database"
	model "portfolio-web-be/models"
	"portfolio-web-be/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateProject(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid form"})
	}

	// Parse descriptions JSON string to []string
	var descriptions []string
	if len(form.Value["descriptions"]) > 0 {
		if err := json.Unmarshal([]byte(form.Value["descriptions"][0]), &descriptions); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid descriptions format"})
		}
	}

	project := model.Project{
		UserID:       userId,
		ProjectName:  form.Value["projectName"][0],
		Descriptions: descriptions,
		TechStack:    form.Value["techStack"][0],
		LinkGithub:   form.Value["linkGithub"][0],
		LinkWebsite:  form.Value["linkWebsite"][0],
	}

	// Handle upload logo 
	if len(form.File["logo"]) > 0 {
		logo := form.File["logo"][0]
		file, err := logo.Open()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to open file"})
		}
		defer file.Close()

		url, err := utils.UploadImageToCloudinary(file, logo)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Upload logo failed"})
		}
		project.Logo = url
	}

	// Save project
	if err := database.DB.Create(&project).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Save project failed"})
	}

	var result model.Project
	if err := database.DB.Preload("User").First(&result, project.ID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch saved project"})
	}

	return c.JSON(result)
}


func GetProject(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)
	var projects []model.Project
	if err := database.DB.Preload("User").Where("user_id = ?", userId).Find(&projects).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Fetch projects failed"})
	}
	return c.JSON(projects)
}