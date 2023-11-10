package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = Init()
	if err != nil {
		log.Fatal("Error initializing database: ", err)
	}

	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})
	app.Use(recover.New())
	app.Use(logger.New())

	setRoutes(app)

	err = app.Listen(":3000")
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
