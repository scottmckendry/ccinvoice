package main

import (
	"log"
	"os"
	"strings"
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
	config := fiber.Config{
		Views: html.New("./views", ".html"),
	}

	// Configure proxy settings only if TRUSTED_PROXIES is set
	if envProxies := os.Getenv("TRUSTED_PROXIES"); envProxies != "" {
		trustedProxies := strings.Split(envProxies, ",")
		for i := range trustedProxies {
			trustedProxies[i] = strings.TrimSpace(trustedProxies[i])
		}

		config.TrustProxy = true
		config.ProxyHeader = fiber.HeaderXForwardedFor
		config.TrustProxyConfig = fiber.TrustProxyConfig{
			Proxies: trustedProxies,
		}
		config.EnableIPValidation = true
	}

	app := fiber.New(config)
	app.Use(recover.New())
	app.Use(logger.New())

	setRoutes(app)

	return app
}
