package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/Shopify/gomail"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"

	"ccinvoice/db"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = db.Init()
	if err != nil {
		log.Fatal("Error initializing database: ", err)
	}

	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})
	app.Use(recover.New())
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		dogs, err := db.GetDogs()
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.Render("index", fiber.Map{
			"dogs": dogs,
		})
	})

	app.Get("/dogs", func(c *fiber.Ctx) error {
		dogs, err := db.GetDogs()
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.Render("dogs", fiber.Map{
			"dogs": dogs,
		})
	})

	app.Get("/dogs/add", func(c *fiber.Ctx) error {
		return c.Render("row-add", nil)
	})

	app.Post("/dogs", func(c *fiber.Ctx) error {
		dog := new(db.Dog)
		if err := c.BodyParser(dog); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		err := db.AddDog(*dog)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		dogs, err := db.GetDogs()
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
		err = db.DeleteDog(id)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		dogs, err := db.GetDogs()
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
		dog, err := db.GetDog(id)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.Render("row-edit", dog)
	})

	app.Put("/dogs/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).SendString(err.Error())
		}
		dog := new(db.Dog)
		dog.ID = id
		if err := c.BodyParser(dog); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		err = db.UpdateDog(*dog)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		dogs, err := db.GetDogs()
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.Render("dogs", fiber.Map{
			"dogs": dogs,
		})
	})

	app.Get("/invoice/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).SendString(err.Error())
		}
		dog, err := db.GetDog(id)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.Render("invoice", dog)
	})

	app.Post("/invoice/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).SendString(err.Error())
		}
		SendInvoice(id)
		dogs, err := db.GetDogs()
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.Render("dogs", fiber.Map{
			"dogs": dogs,
		})
	})

	app.Listen(":3000")
}

func SendInvoice(id int) {

	log.Println("Creating PDF")
	pdfGenerator, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal("Error creating PDF generator: ", err)
	}

	page := wkhtmltopdf.NewPage("localhost:3000/invoice/" + strconv.Itoa(id))

	pdfGenerator.AddPage(page)

	err = pdfGenerator.Create()
	if err != nil {
		log.Fatal("Error generating PDF: ", err)
	}

	err = pdfGenerator.WriteFile("./test.pdf")
	if err != nil {
		log.Fatal("Error writing PDF: ", err)
	}

	smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		log.Fatal("Error converting SMTP_PORT to int: ", err)
	}

	log.Println("Sending email")
	d := gomail.NewDialer(
		os.Getenv("SMTP_HOST"),
		smtpPort,
		os.Getenv("SMTP_USER"),
		os.Getenv("SMTP_PASS"),
	)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage()
	m.SetHeader("From", fmt.Sprintf("Test <%s>", os.Getenv("SMTP_USER")))
	m.SetHeader("To", fmt.Sprintf("Test <%s>", os.Getenv("SMTP_USER")))
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>World</b>!")
	m.Attach("./test.pdf")

	err = d.DialAndSend(m)
	if err != nil {
		log.Fatal("Error sending email: ", err)
	} else {
		log.Println("Email sent!")
	}

	log.Println("Deleting PDF")
	err = os.Remove("./test.pdf")
	if err != nil {
		log.Fatal("Error deleting PDF: ", err)
	} else {
		log.Println("PDF deleted!")
	}
}
