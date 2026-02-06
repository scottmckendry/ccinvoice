package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
)

// Renders the index page.
func renderIndex(c fiber.Ctx) error {
	dogs, err := getDogs()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Render("index", fiber.Map{
		"dogs": dogs,
	})
}

// Renders the cards for all dogs.
func renderDogs(c fiber.Ctx) error {
	dogs, err := getDogs()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Render("dogs", fiber.Map{
		"dogs": dogs,
	})
}

// Renders the modal for adding a new dog.
func renderAdd(c fiber.Ctx) error {
	return c.Render("modal-add", nil)
}

// Gets details for the given dog ID and renders the modal for editing.
func renderEdit(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	dog, err := getDog(id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Render("modal-edit", dog)
}

// Generates a preview of the invoice for a given dog as HTML. This is used for generating the PDF.
func renderInvoice(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	dog, err := getDog(id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	total := 0.0
	for _, service := range dog.Services {
		total += float64(service.Quantity) * service.Price
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
		"Services":      dog.Services,
		"Total":         strconv.FormatFloat(total, 'f', 2, 64),
	})
}

// Generates a PDF version of the HTML returned by renderInvoice() using the wkhtmltopdf command line tool.
func renderInvoicePdf(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	dog, err := getDog(id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	pdfPath, err := generatePdf(dog)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.SendFile(pdfPath)
}

// Adds a new dog to the database and returns the updated list of dogs.
func handleDogAdd(c fiber.Ctx) error {
	dog := new(Dog)
	if err := c.Bind().Body(dog); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err := addDog(*dog)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	dogs, err := getDogs()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Render("dogs", fiber.Map{
		"dogs": dogs,
	})
}

// Updates a dog matching the provided ID in the database and returns the updated list of dogs.
func handleDogUpdate(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	dog := new(Dog)
	dog.ID = id
	if err := c.Bind().Body(dog); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err = updateDog(*dog)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	dogs, err := getDogs()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Render("dogs", fiber.Map{
		"dogs": dogs,
	})
}

// Deletes a dog matching the provided ID from the database and returns the updated list of dogs.
func handleDogDelete(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err = deleteDog(id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	dogs, err := getDogs()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Render("dogs", fiber.Map{
		"dogs": dogs,
	})
}

// Sends an invoice for the given dog to the owner's email address.
func handleSendInvoice(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	// err = sendInvoice(id)
	// err = nil // replace with db insert into email queue table
	err = queueEmail(id)
	if err != nil {
		log.Println(err)
		// We're intentionally not returning the error to the client here. HTMX won't update the element if the response is an error.
		// Not sure why this is the case, but probably worth revisiting in the future.
		return c.SendString("Failed :(")
	}
	return c.SendString("<button class=\"btn-disabled\" disabled>Queued!</button>")
}
