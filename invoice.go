package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Shopify/gomail"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func sendInvoice(id int) error {
	dog, err := getDog(id)
	if err != nil {
		return err
	}

	_, err = generatePdf(dog)
	if err != nil {
		return err
	}

	err = sendEmail(dog)
	if err != nil {
		return err
	}

	return nil
}

func sendInvoices() (status string, err error) {
	emails, err := getEmailQueue()
	if err != nil {
		return "", err
	}

	err = markEmailsInProcess(emails)
	if err != nil {
		return "", err
	}

	for _, email := range emails {
		if err != nil {
			return "", err
		}
		err = sendInvoice(email.DogID)
		if err != nil {
			return "", err
		}

		err = markEmailSent(email.ID)
		if err != nil {
			return "", err
		}
	}

	return "Processed " + strconv.Itoa(len(emails)) + " emails", nil
}

func generatePdf(dog Dog) (string, error) {
	// Create a headless Chrome context with no sandbox - required for running in non-root environments
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.NoSandbox,
		chromedp.Flag("disable-setuid-sandbox", true),
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var pdfBuf []byte
	err := chromedp.Run(ctx,
		chromedp.Navigate(os.Getenv("BASE_URL")+"/invoice/"+strconv.Itoa(dog.ID)),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			pdfBuf, _, err = page.PrintToPDF().
				WithPrintBackground(true).
				WithScale(0.8).
				WithPaperHeight(12).
				Do(ctx)
			return err
		}),
	)

	if err != nil {
		return "", fmt.Errorf("error generating PDF: %w", err)
	}
	invoiceFile := fmt.Sprintf("./public/%s.pdf", getInvoiceNumber(dog))
	// Save to file
	if err := os.WriteFile(invoiceFile, pdfBuf, 0644); err != nil {
		return "", fmt.Errorf("error writing PDF to file: %w", err)
	}
	return invoiceFile, nil
}

func sendEmail(dog Dog) error {
	invoiceFile := fmt.Sprintf("./public/%s.pdf", getInvoiceNumber(dog))
	smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		return fmt.Errorf("error converting SMTP_PORT to int: %s", err)
	}

	d := gomail.NewDialer(
		os.Getenv("SMTP_HOST"),
		smtpPort,
		os.Getenv("SMTP_USER"),
		os.Getenv("SMTP_PASS"),
	)
	d.Timeout = 30 * time.Second
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	ownerFirstName := strings.Split(dog.OwnerName, " ")[0]
	fromFirstName := strings.Split(os.Getenv("FROM_NAME"), " ")[0]

	m := gomail.NewMessage()
	m.SetHeader("From", fmt.Sprintf("Canine Club<%s>", os.Getenv("SMTP_USER")))
	m.SetHeader("To", fmt.Sprintf("%s <%s>", dog.OwnerName, dog.Email))
	m.SetHeader("Subject", "Canine Club - Invoice for "+dog.Name)
	m.SetBody(
		"text/html",
		"Hi "+ownerFirstName+",<br><br>Please find attached the invoice for "+dog.Name+"'s walks this week.<p style='font-weight:lighter;'>Please use '<b>"+dog.Name+"</b>' as the reference when making payment. Also note that payment is due by "+nextMonday(
			time.Now(),
		).Format("Monday, 2 January 2006")+
			".</p><br>Any questions let me know,<br>Thank you!<br><br>"+fromFirstName+"<br>Canine Club",
	)
	m.Attach(invoiceFile)

	err = d.DialAndSend(m)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func getInvoiceNumber(dog Dog) string {
	name := dog.Name
	if len(name) < 3 {
		name = name + strings.Repeat("0", 3-len(name))
	}
	prefix := strings.ToUpper(name[0:3])
	return prefix + time.Now().Format("20060102")
}

func nextMonday(t time.Time) time.Time {
	if t.Weekday() == time.Monday {
		return t.AddDate(0, 0, 7)
	}
	for t.Weekday() != time.Monday {
		t = t.AddDate(0, 0, 1)
	}
	return t
}
