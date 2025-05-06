package main

import (
	"os"
	"portfolio-web-be/config"
	"portfolio-web-be/database"
	"portfolio-web-be/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	database.Connect()

	
	app := fiber.New()

	routes.LoginRoutes(app)

	port := os.Getenv("PORT")
	
	app.Listen(":" + port)
}