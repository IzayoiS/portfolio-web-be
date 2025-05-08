package service

import (
	"portfolio-web-be/database"
	model "portfolio-web-be/models"
	"portfolio-web-be/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateExperience(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid form data"})
	}

	exp := model.Experience{
		UserID: 		   userId,
		Company: 		   form.Value["company"][0],
		Role:              form.Value["role"][0],
		StartMonth:        form.Value["startMonth"][0],
		StartYear:         form.Value["startYear"][0],
		IsCurrentlyWorking: form.Value["isCurrentlyWorking"][0] == "true",
		Descriptions:      form.Value["descriptions"],
	}

	if len(form.Value["endMonth"]) > 0 {
		exp.EndMonth = &form.Value["endMonth"][0]
	}
	if len(form.Value["endYear"]) > 0 {
		exp.EndYear = &form.Value["endYear"][0]
	}

	if len(form.File["logo"]) > 0 {
		logoFile := form.File["logo"][0]
		file, err := logoFile.Open()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to open logo file"})
		}
		defer file.Close()

		url, err := utils.UploadImageToCloudinary(file, logoFile)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to upload logo to Cloudinary"})
		}
		exp.Logo = url
	}

	if err := database.DB.Create(&exp).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save experience"})
	}

	if err := database.DB.Preload("User").First(&exp, exp.ID).Error; err != nil {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": "Failed to fetch created experience with user",
	})
}
	return c.JSON(exp)
}

func GetExperience(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)
	var experiences []model.Experience
	if err := database.DB.Preload("User").Where("user_id = ?", userId).Find(&experiences).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not fetch experiences"})
	}
	return c.JSON(experiences)
}