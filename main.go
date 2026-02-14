package main

import (
	"log"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/template/html/v3"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Could not load .env file: %v | Continuing with system variables", err)
	}

	err = Init()
	if err != nil {
		log.Fatal("Error initializing database: ", err)
	}

	err = startScheduler()
	if err != nil {
		log.Fatal("Error starting scheduler: ", err)
	}

	app := startServer()
	err = app.Listen(":3000")
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

func startScheduler() error {
	s, err := gocron.NewScheduler()
	if err != nil {
		return err
	}

	j, err := s.NewJob(
		gocron.DurationJob(10*time.Second),
		gocron.NewTask(func() {
			status, err := sendInvoices()
			if err != nil {
				log.Println("Error sending invoices: ", err)
			}
			if status != "Processed 0 emails" {
				log.Println(status)
			}
		}),
	)

	if err != nil {
		return err
	}

	log.Println("Job ID: ", j.ID())
	s.Start()

	return nil
}

func startServer() *fiber.App {
	app := fiber.New(fiber.Config{
		TrustProxy: true,
		TrustProxyConfig: fiber.TrustProxyConfig{
			Private:   true,
			LinkLocal: true,
			Loopback:  true,
		},
		Views: html.New("./views", ".html"),
	})
	app.Use(recover.New())
	app.Use(logger.New())

	setRoutes(app)

	return app
}
