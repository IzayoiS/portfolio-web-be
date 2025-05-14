package controller

import (
	model "portfolio-web-be/models"
	service "portfolio-web-be/services"
	"portfolio-web-be/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateProfile(c *fiber.Ctx) error {
	var profile model.Profile

	if err := c.BodyParser(&profile); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	userId := c.Locals("userId").(uint)
	profile.UserID = userId
	
	result, err := service.CreateProfile(&profile)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create profile"})
	}
	return c.JSON(result)
}

func GetProfile(c *fiber.Ctx) error {
	userID := c.Params("userId")
	var profile *model.Profile

	profile, err := service.GetProfileByUserID(utils.ParseUint(userID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Profile not found"})
	}
	return c.JSON(profile)
}

func UpdateProfile(c *fiber.Ctx) error {
	userID := c.Params("userId")

	// Fetch data old profile from DB
	oldProfile, err := service.GetProfileByUserID(utils.ParseUint(userID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Profile not found"})
	}

	var oldImagePublicID string
	if oldProfile.ImageURL != "" {
		oldImagePublicID = utils.ExtractPublicIDFromURL(oldProfile.ImageURL)
	}

	// Fetch data from form-data
	updatedProfile := model.Profile{
		Name:         c.FormValue("name"),
		JobTitle:     c.FormValue("job_title"),
		Bio:          c.FormValue("bio"),
		Location:     c.FormValue("location"),
		Availability: c.FormValue("availability"),
	}

	fileHeader, err := c.FormFile("image")
	if err == nil {
		file, err := fileHeader.Open()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to open image file"})
		}
		defer file.Close()

		// Upload new image 
		imageUrl, err := utils.UploadImageToCloudinary(file, fileHeader)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to upload image"})
		}

		// Delete old image from cloudinary 
		if oldImagePublicID != "" {
			if err := utils.DeleteImage(oldImagePublicID); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete old image"})
			}
		}

		updatedProfile.ImageURL = imageUrl
	} else {
		updatedProfile.ImageURL = oldProfile.ImageURL
	}

	// Save to database
	updatedProfileData, err := service.UpdateProfile(utils.ParseUint(userID), &updatedProfile)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update profile"})
	}

	return c.JSON(updatedProfileData)
}


func UploadProfileImage(c *fiber.Ctx) error {
	fileHeader, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Image is required"})
	}

	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	imageUrl, err := utils.UploadImageToCloudinary(file,fileHeader)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to upload image"})
	}

	return c.JSON(fiber.Map{"image_url": imageUrl})
}