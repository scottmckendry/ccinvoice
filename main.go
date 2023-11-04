package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/Shopify/gomail"
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

	app.Get("/", func(c *fiber.Ctx) error {
		dogs, err := GetDogs()
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.Render("index", fiber.Map{
			"dogs": dogs,
		})
	})

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
		return c.Render("row-add", nil)
	})

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
		return c.Render("row-edit", dog)
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
			"DueDate":       nextMonday().Format("Monday, 2 January 2006"),
			"Name":          dog.Name,
			"OwnerName":     dog.OwnerName,
			"Address":       dog.Address,
			"City":          dog.City,
			"Walks":         dog.WalksPerWeek,
			"Price":         strconv.FormatFloat(dog.PricePerWalk, 'f', 2, 64),
			"Total": strconv.FormatFloat(
				(float64(dog.WalksPerWeek) * dog.PricePerWalk),
				'f',
				2,
				64,
			),
		})
	})

	app.Post("/invoice/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).SendString(err.Error())
		}
		SendInvoice(id)
		dogs, err := GetDogs()
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
	dog, err := GetDog(id)
	if err != nil {
		log.Fatal("Error getting dog: ", err)
	}

	log.Println("Creating PDF")
	pdfGenerator, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal("Error creating PDF generator: ", err)
	}

	page := wkhtmltopdf.NewPage("http://localhost:3000/invoice/" + strconv.Itoa(id))

	pdfGenerator.AddPage(page)

	err = pdfGenerator.Create()
	if err != nil {
		log.Fatal("Error generating PDF: ", err)
	}

	invoiceFile := fmt.Sprintf("./%s.pdf", getInvoiceNumber(dog))

	err = pdfGenerator.WriteFile(invoiceFile)
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
	m.SetHeader("From", fmt.Sprintf("Canine Club<%s>", os.Getenv("SMTP_USER")))
	m.SetHeader("To", fmt.Sprintf("%s <%s>", dog.OwnerName, os.Getenv("SMTP_USER")))
	m.SetHeader("Subject", "Canine Club - Invoice for "+dog.Name)
	m.SetBody(
		"text/html",
		"Hi "+dog.OwnerName+",<br><br>Here is your invoice for "+dog.Name+".<br><br>Kind regards,<br>Canine Club",
	)
	m.Attach(invoiceFile)

	err = d.DialAndSend(m)
	if err != nil {
		log.Fatal("Error sending email: ", err)
	} else {
		log.Println("Email sent!")
	}

	log.Println("Deleting PDF")
	err = os.Remove(invoiceFile)
	if err != nil {
		log.Fatal("Error deleting PDF: ", err)
	} else {
		log.Println("PDF deleted!")
	}
}

func getInvoiceNumber(dog Dog) string {
	prefix := strings.ToUpper(dog.Name[0:3])
	return prefix + time.Now().Format("20060102")
}

func nextMonday() time.Time {
	today := time.Now()
	daysUntilMonday := int(time.Monday - today.Weekday())
	if daysUntilMonday < 0 {
		daysUntilMonday += 7
	}
	return today.AddDate(0, 0, daysUntilMonday)
}
