package main

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {
	// Public files
	app.Static("/", "./public")

	// Renders the index view with a list of all dogs.
	app.Get("/", func(c *fiber.Ctx) error {
		dogs, err := GetDogs()
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.Render("index", fiber.Map{
			"dogs": dogs,
		})
	})

	// Renders the dogs view with a list of all dogs.
	app.Get("/dogs", func(c *fiber.Ctx) error {
		dogs, err := GetDogs()
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.Render("dogs", fiber.Map{
			"dogs": dogs,
		})
	})

	app.Get("/dogs/add", func(c *fiber.Ctx) error {
		return c.Render("modal-add", nil)
	})

	// Adds a new dog to the database and returns the updated list of dogs.
	app.Post("/dogs", func(c *fiber.Ctx) error {
		dog := new(Dog)
		if err := c.BodyParser(dog); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		err := AddDog(*dog)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		dogs, err := GetDogs()
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.Render("dogs", fiber.Map{
			"dogs": dogs,
		})
	})

	app.Delete("/dogs/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).SendString(err.Error())
		}
		err = DeleteDog(id)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		dogs, err := GetDogs()
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.Render("dogs", fiber.Map{
			"dogs": dogs,
		})
	})

	app.Get("dogs/edit/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).SendString(err.Error())
		}
		dog, err := GetDog(id)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.Render("modal-edit", dog)
	})

	app.Put("/dogs/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).SendString(err.Error())
		}
		dog := new(Dog)
		dog.ID = id
		if err := c.BodyParser(dog); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		err = UpdateDog(*dog)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		dogs, err := GetDogs()
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.Render("dogs", fiber.Map{
			"dogs": dogs,
		})
	})

	// Generates a preview of the invoice for a given dog.
	// Intended to be opened in a new tab.
	app.Get("/invoice/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).SendString(err.Error())
		}
		dog, err := GetDog(id)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.Render("invoice", fiber.Map{
			"InvoiceNumber": getInvoiceNumber(dog),
			"Date":          time.Now().Format("Monday, 2 January 2006"),
			"DueDate":       nextMonday(time.Now()).Format("Monday, 2 January 2006"),
			"FromName":      os.Getenv("FROM_NAME"),
			"FromAddress":   os.Getenv("FROM_ADDRESS"),
			"FromCity":      os.Getenv("FROM_CITY"),
			"AccountNumber": os.Getenv("ACCOUNT_NUMBER"),
			"Name":          dog.Name,
			"OwnerName":     dog.OwnerName,
			"Address":       dog.Address,
			"City":          dog.City,
			"Service":       dog.Service,
			"Quantity":      dog.Quantity,
			"Price":         strconv.FormatFloat(dog.Price, 'f', 2, 64),
			"Total": strconv.FormatFloat(
				(float64(dog.Quantity) * dog.Price),
				'f',
				2,
				64,
			),
		})
	})

	// Returns a PDF of the invoice for a given dog.
	app.Get("/invoice/:id/pdf", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).SendString(err.Error())
		}
		dog, err := GetDog(id)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		pdfPath, err := generatePdf(dog)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.SendFile(pdfPath)
	})

	// Generates a PDF of the invoice for a given dog and emails it to the owner.
	app.Post("/invoice/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).SendString(err.Error())
		}
		err = sendInvoice(id)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.SendString("Done!")
	})
}
