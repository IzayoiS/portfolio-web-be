package service

import (
	// "encoding/json"
	"encoding/json"
	"portfolio-web-be/database"
	model "portfolio-web-be/models"
	"portfolio-web-be/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

func CreateExperience(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid form data"})
	}

	var descriptions []string
	descStr := getFormValue(form, "descriptions")
	if descStr != "" {
		if err := json.Unmarshal([]byte(descStr), &descriptions); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid descriptions JSON"})
		}
	}


	isCurrentlyWorking := form.Value["isCurrentlyWorking"][0] == "true"
	exp := model.Experience{
		UserID: 		   userId,
		Company: 		   form.Value["company"][0],
		Role:              form.Value["role"][0],
		StartMonth:        form.Value["startMonth"][0],
		StartYear:         form.Value["startYear"][0],
		IsCurrentlyWorking: &isCurrentlyWorking,
		Descriptions:      descriptions,
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

func UpdateExperience(c *fiber.Ctx) error {
	expId := c.Params(("id"))
	userId := c.Locals("userId").(uint)

	id, err := strconv.ParseUint(expId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid experience ID"})
	}

	var existingExp model.Experience
	if err := database.DB.Where("id = ? AND user_id = ?", uint(id), userId).First(&existingExp).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Experience not found or unauthorized"})
	}

	form ,err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid form"})
	}

    var endMonth *string
    if val := getFormValue(form, "end_month"); val != "" {
        endMonth = &val
    }

    var endYear *string
    if val := getFormValue(form, "end_year"); val != "" {
        endYear = &val
    }

	var isCurrentlyWorking *bool
	if val := getFormValue(form, "currently_working"); val != "" {
		b := val == "true"
		isCurrentlyWorking = &b
	}
	var descriptions []string
	descStr := getFormValue(form, "descriptions")
	if descStr != "" {
		if err := json.Unmarshal([]byte(descStr), &descriptions); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid descriptions JSON"})
		}
	}

	
	updates := model.Experience{
        Company:  getFormValue(form, "company"),
        Role: getFormValue(form, "role"),
        StartMonth:    getFormValue(form, "start_month"),
        StartYear:   getFormValue(form, "start_year"),
		EndMonth:           endMonth,
        EndYear:            endYear,
        IsCurrentlyWorking: isCurrentlyWorking,
		Descriptions:       pq.StringArray(descriptions),
    }
	
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
        updates.Logo = url
        
        if existingExp.Logo != "" {
            publicID := utils.ExtractPublicIDFromURL(existingExp.Logo)
            if publicID != "" {
                utils.DeleteImage(publicID)
            }
        }
    } else if logoUrl := getFormValue(form, "logo_url"); logoUrl != "" {
        updates.Logo = logoUrl
    }

	if err := database.DB.Model(&existingExp).Updates(updates).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Update experience failed"})
    }

	var result model.Experience
    if err := database.DB.Preload("User").First(&result, existingExp.ID).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch updated experience"})
    }

    return c.JSON(result)
}

func DeleteExperience(c *fiber.Ctx) error {
	expId := c.Params("id")
	userId := c.Locals("userId").(uint)

	var exp model.Experience
    if err := database.DB.Where("id = ? AND user_id = ?", expId, userId).First(&exp).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Experience not found"})
    }

	if exp.Logo != "" {
        publicID := utils.ExtractPublicIDFromURL(exp.Logo)
        if publicID != "" {
            if err := utils.DeleteImage(publicID); err != nil {
                return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete logo"})
            }
        }
    }

    if err := database.DB.Delete(&exp).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Delete experience failed"})
    }

    return c.JSON(fiber.Map{"message": "Experience deleted successfully"})
}