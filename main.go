package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Shopify/gomail"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	log.Println("Loading .env file")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
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
	m.Attach("./test.txt")

	err = d.DialAndSend(m)
	if err != nil {
		log.Fatal("Error sending email: ", err)
	} else {
		log.Println("Email sent!")
	}
}
