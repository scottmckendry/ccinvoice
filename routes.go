package main

import (
	"github.com/gofiber/fiber/v2"
)

func setRoutes(app *fiber.App) {
	// Serve static files from the public directory
	app.Static("/", "./public")

	// Render templates
	app.Get("/", renderIndex)
	app.Get("/dogs", renderDogs)
	app.Get("/dogs/add", renderAdd)
	app.Get("dogs/edit/:id", renderEdit)
	app.Get("/invoice/:id", renderInvoice)
	app.Get("/invoice/:id/pdf", renderInvoicePdf)

	// API endpoints
	app.Post("/dogs", handleDogAdd)
	app.Put("/dogs/:id", handleDogUpdate)
	app.Delete("/dogs/:id", handleDogDelete)
	app.Post("/invoice/:id", handleSendInvoice)
}
