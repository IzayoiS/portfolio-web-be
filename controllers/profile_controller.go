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

	// Ambil data profile lama dari DB
	oldProfile, err := service.GetProfileByUserID(utils.ParseUint(userID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Profile not found"})
	}

	var oldImagePublicID string
	if oldProfile.ImageURL != "" {
		oldImagePublicID = oldProfile.ImageURL
	}

	// Ambil data dari form (bukan BodyParser karena ini multipart/form-data)
	updatedProfile := model.Profile{
		Name:         c.FormValue("name"),
		JobTitle:     c.FormValue("job_title"),
		Bio:          c.FormValue("bio"),
		Location:     c.FormValue("location"),
		Availability: c.FormValue("availability"),
	}

	// Proses upload gambar jika ada
	fileHeader, err := c.FormFile("image")
	if err == nil {
		file, err := fileHeader.Open()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to open image file"})
		}
		defer file.Close()

		// Upload ke Cloudinary
		imageUrl, err := utils.UploadImageToCloudinary(file, fileHeader)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to upload image"})
		}

		// Hapus gambar lama jika ada
		if oldImagePublicID != "" {
			if err := utils.DeleteImage(oldImagePublicID); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete old image"})
			}
		}

		updatedProfile.ImageURL = imageUrl
	} else {
		// Jika tidak ada gambar baru, gunakan yang lama
		updatedProfile.ImageURL = oldProfile.ImageURL
	}

	// Update profile ke DB
	updatedProfileData, err := service.UpdateProfile(utils.ParseUint(userID), &updatedProfile, oldImagePublicID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update profile"})
	}

	return c.JSON(updatedProfileData)
}

// func UpdateProfile(c *fiber.Ctx) error {
// 	userID := c.Params("userId")

// 	var updatedProfile model.Profile
// 	if err := c.BodyParser(&updatedProfile); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
// 	}

// 	oldProfile, err := service.GetProfileByUserID(utils.ParseUint(userID))
// 	if err != nil {
// 		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Profile not found"})
// 	}

// 	var oldImagePublicID string
// 	if oldProfile.ImageURL != "" {
// 		oldImagePublicID = oldProfile.ImageURL
// 	}
	
// 	fileHeader, err := c.FormFile("image")
// 	if err == nil {
// 		file, err := fileHeader.Open()
// 		if err != nil {
// 			return err
// 		}
// 		defer file.Close()

// 		imageUrl, err := utils.UploadImageToCloudinary(file, fileHeader)
// 		if err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to upload image"})
// 		}

// 		if oldImagePublicID != "" {
// 			if err := utils.DeleteImage(oldImagePublicID); err != nil {
// 				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete old image"})
// 			}
// 		}
// 		updatedProfile.ImageURL = imageUrl
// 	}
// 	updatedProfileData, err := service.UpdateProfile(utils.ParseUint(userID), &updatedProfile, oldImagePublicID)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update profile"})
// 	}

// 	return c.JSON(updatedProfileData)
// }


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