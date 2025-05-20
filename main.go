package main

import (
	"log"
	"os"
	"portfolio-web-be/config"
	"portfolio-web-be/database"
	model "portfolio-web-be/models"
	"portfolio-web-be/routes"

	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	database.Connect()

	if err := database.DB.AutoMigrate(&model.User{}, &model.Profile{},&model.Experience{},&model.Project{},&model.TechStack{}); err != nil {
        log.Fatal("AutoMigrate failed:", err)
    }
	log.Println("Database migrated successfully!")
	
	app := fiber.New()
	app.Use(cors.New(cors.Config{
        AllowOrigins: "http://localhost:3000, https://portfolio-iqbals.vercel.app/", 
        AllowHeaders: "Origin, Content-Type, Accept, Authorization",
    }))

	routes.LoginRoutes(app)
	routes.ProfileRoute(app)
	routes.ExperienceRoutes(app)
	routes.ProjectRoute(app)
	routes.TechRoutes(app)
	routes.RegisterRoute(app)
	routes.CheckUserRoutes(app)

	port := os.Getenv("PORT")
	
	app.Listen(":" + port)
}