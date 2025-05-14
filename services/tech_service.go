package service

import (
	"portfolio-web-be/database"
	model "portfolio-web-be/models"
	"portfolio-web-be/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateTechStack(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)

	// Ambil nama dari form-data
	name := c.FormValue("name")
	if name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Name is required"})
	}

	// Ambil file icon dari form-data
	fileHeader, err := c.FormFile("icon")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Icon file is required"})
	}

	file, err := fileHeader.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to open icon file"})
	}
	defer file.Close()

	// Upload ke Cloudinary
	iconURL, err := utils.UploadImageToCloudinary(file, fileHeader)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to upload image to Cloudinary"})
	}

	// Simpan ke DB
	tech := model.TechStack{
		UserID: userId,
		Name:   name,
		Icon:   iconURL,
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

func UpdateTech(c *fiber.Ctx) error {
	techId := c.Params("id")
	userId := c.Locals("userId").(uint)

	id, err := strconv.ParseUint(techId, 10, 32)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid project ID"})
    }

	var existingTech model.TechStack
	if err := database.DB.Where("id = ? AND user_id = ?", uint(id),userId).First(&existingTech).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tech not found or unauthorized"})
	}
	
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid form"})
	}

	updates := model.TechStack{
		Name: getFormValue(form, "name"),
	}

	if len(form.File["icon"]) > 0 {
		icon := form.File["icon"][0]
		file, err := icon.Open()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to open file"})
		}
		defer file.Close()

		url, err := utils.UploadImageToCloudinary(file, icon)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Upload icon failed"})
		}
		updates.Icon = url

		if existingTech.Icon != "" {
			publicID := utils.ExtractPublicIDFromURL(existingTech.Icon)
			if publicID != "" {
				utils.DeleteImage(publicID)
			}
		}
	} else if iconUrl := getFormValue(form, "icon_url"); iconUrl != "" {
		updates.Icon = iconUrl
	}

	if err := database.DB.Model(&existingTech).Updates(updates).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Update project failed"})
	}

	var result model.TechStack
	if err := database.DB.Preload("User").First(&result, existingTech.ID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch updated project"})
	}

	return c.JSON(result)
}

func DeleteTech(c *fiber.Ctx) error {
	techId := c.Params("id")
	userId := c.Locals("userId").(uint)

	var tech model.TechStack
	if err := database.DB.Where("id = ? AND user_id = ?", techId, userId).First(&tech).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Project not found"})
	}

	if tech.Icon != "" {
		publicID := utils.ExtractPublicIDFromURL(tech.Icon)
		if publicID != "" {
			if err := utils.DeleteImage(publicID); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete icon"})
			}
		}
	}

	if err := database.DB.Delete(&tech).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Delete tech failed"})
	}

	return c.JSON(fiber.Map{"message": "Tech deleted successfully"})
}