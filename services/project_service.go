package service

import (
	"mime/multipart"
	"portfolio-web-be/database"
	model "portfolio-web-be/models"
	"portfolio-web-be/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func getFormValue(form *multipart.Form, key string) string {
    if values, ok := form.Value[key]; ok && len(values) > 0 {
        return values[0]
    }
    return ""
}

func CreateProject(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid form"})
	}

	project := model.Project{
		UserID:       userId,
		ProjectName:  getFormValue(form, "project_name"),
		Descriptions: getFormValue(form, "descriptions"),
		TechStack:    getFormValue(form, "tech_stack"),
		LinkGithub:   getFormValue(form, "link_github"),
		LinkWebsite:  getFormValue(form, "link_website"),
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
		project.Logo = url
	}

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

func UpdateProject(c *fiber.Ctx) error {
    projectId := c.Params("id")
    userId := c.Locals("userId").(uint)
    
	id, err := strconv.ParseUint(projectId, 10, 32)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid project ID"})
    }
	
    var existingProject model.Project
    if err := database.DB.Where("id = ? AND user_id = ?", uint(id), userId).First(&existingProject).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Project not found or unauthorized",
            "details": err.Error(), 
        })
    }


    form, err := c.MultipartForm()
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid form"})
    }

    updates := model.Project{
        ProjectName:  getFormValue(form, "project_name"),
        Descriptions: getFormValue(form, "descriptions"),
        TechStack:    getFormValue(form, "tech_stack"),
        LinkGithub:   getFormValue(form, "link_github"),
        LinkWebsite:  getFormValue(form, "link_website"),
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
        
        if existingProject.Logo != "" {
            publicID := utils.ExtractPublicIDFromURL(existingProject.Logo)
            if publicID != "" {
                utils.DeleteImage(publicID)
            }
        }
    } else if logoUrl := getFormValue(form, "logo_url"); logoUrl != "" {
        updates.Logo = logoUrl
    }

    if err := database.DB.Model(&existingProject).Updates(updates).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Update project failed"})
    }

    var result model.Project
    if err := database.DB.Preload("User").First(&result, existingProject.ID).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch updated project"})
    }

    return c.JSON(result)
}

func DeleteProject(c *fiber.Ctx) error {
    projectId := c.Params("id")
    userId := c.Locals("userId").(uint)

    var project model.Project
    if err := database.DB.Where("id = ? AND user_id = ?", projectId, userId).First(&project).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Project not found"})
    }

    if project.Logo != "" {
        publicID := utils.ExtractPublicIDFromURL(project.Logo)
        if publicID != "" {
            if err := utils.DeleteImage(publicID); err != nil {
                return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete logo"})
            }
        }
    }

    if err := database.DB.Delete(&project).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Delete project failed"})
    }

    return c.JSON(fiber.Map{"message": "Project deleted successfully"})
}