package main

import (
	"io"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
)

var app *fiber.App

func TestViews(t *testing.T) {
	app = fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})
	app.Use(recover.New())
	setRoutes(app)

	routes := []string{"/", "/dogs", "/dogs/add", "/dogs/edit/1", "/invoice/1", "/invoice/1/pdf"}

	for _, route := range routes {
		req := httptest.NewRequest("GET", route, nil)
		resp, err := app.Test(req, 10000)
		if err != nil {
			t.Error("Error sending request to Fiber: ", err)
		}

		if resp.StatusCode != 200 {
			t.Error("Expected status code 200, got ", resp.StatusCode)
			t.Error("Route: ", route)
			bodyString, _ := io.ReadAll(resp.Body)
			t.Error("Body: ", string(bodyString))
		}
	}

	// Test that a 500 error is returned if there is an error rendering the template
	oldDbUrl := dbUrl
	dbUrl = "someBadUrl"
	connect()

	for _, route := range routes {
		if route == "/dogs/add" {
			continue // This route has no database interaction
		}

		req := httptest.NewRequest("GET", route, nil)
		resp, err := app.Test(req, 10000)
		if err != nil {
			t.Error("Error sending request to Fiber: ", err)
		}

		if resp.StatusCode != 500 {
			t.Error("Expected status code 500, got ", resp.StatusCode)
			t.Error("Expected status code 200, got ", resp.StatusCode)
			t.Error("Route: ", route)
			bodyString, _ := io.ReadAll(resp.Body)
			t.Error("Body: ", string(bodyString))
		}

	}

	// Reverse the changes to the database connection
	t.Cleanup(func() {
		dbUrl = oldDbUrl
		Init()
	})
}

func TestBadPaths(t *testing.T) {
	app = fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})
	app.Use(recover.New())
	setRoutes(app)

	routes := []string{
		"/badpath",
		"/dogs/badpath",
		"/dogs/edit/badpath",
		"/invoice/badpath",
		"/invoice/badpath/pdf",
	}

	for _, route := range routes {
		req := httptest.NewRequest("GET", route, nil)
		resp, err := app.Test(req, 10000)
		if err != nil {
			t.Error("Error sending request to Fiber: ", err)
		}

		if resp.StatusCode < 400 || resp.StatusCode > 405 {
			t.Error("Expected status code 40X, got ", resp.StatusCode)
			t.Error("Route: ", route)
			bodyString, _ := io.ReadAll(resp.Body)
			t.Error("Body: ", string(bodyString))
		}
	}
}

func TestPostReq(t *testing.T) {
	app = fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})
	app.Use(recover.New())
	setRoutes(app)

	formData := url.Values{
		"name":      {"Rex"},
		"ownerName": {"Jane Doe"},
	}

	req := httptest.NewRequest("POST", "/dogs", strings.NewReader(formData.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := app.Test(req, 10000)
	if err != nil {
		t.Error("Error sending request to Fiber: ", err)
	}

	if resp.StatusCode != 200 {
		t.Error("Expected status code 200, got ", resp.StatusCode)
		bodyString, _ := io.ReadAll(resp.Body)
		t.Error("Body: ", string(bodyString))
	}

	formData = url.Values{
		"name":      {"Rex"},
		"ownerName": {"Jane Doe"},
		"price":     {"not a number"},
	}

	req = httptest.NewRequest("POST", "/dogs", strings.NewReader(formData.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err = app.Test(req, 10000)
	if err != nil {
		t.Error("Error sending request to Fiber: ", err)
	}

	if resp.StatusCode != 400 {
		t.Error("Expected status code 400, got ", resp.StatusCode)
		bodyString, _ := io.ReadAll(resp.Body)
		t.Error("Body: ", string(bodyString))
	}
}

func TestPutReq(t *testing.T) {
	app = fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})
	app.Use(recover.New())
	setRoutes(app)

	formData := url.Values{
		"name":      {"Ralph"},
		"ownerName": {"James Doe"},
	}

	req := httptest.NewRequest("PUT", "/dogs/1", strings.NewReader(formData.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := app.Test(req, 10000)
	if err != nil {
		t.Error("Error sending request to Fiber: ", err)
	}

	if resp.StatusCode != 200 {
		t.Error("Expected status code 200, got ", resp.StatusCode)
		bodyString, _ := io.ReadAll(resp.Body)
		t.Error("Body: ", string(bodyString))
	}

	formData = url.Values{
		"name":      {"Ralph"},
		"ownerName": {"James Doe"},
		"price":     {"not a number"},
	}

	req = httptest.NewRequest("PUT", "/dogs/1", strings.NewReader(formData.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err = app.Test(req, 10000)
	if err != nil {
		t.Error("Error sending request to Fiber: ", err)
	}

	if resp.StatusCode != 400 {
		t.Error("Expected status code 400, got ", resp.StatusCode)
		bodyString, _ := io.ReadAll(resp.Body)
		t.Error("Body: ", string(bodyString))
	}

	req = httptest.NewRequest("PUT", "/dogs/abc", nil)
	resp, err = app.Test(req, 10000)
	if err != nil {
		t.Error("Error sending request to Fiber: ", err)
	}

	if resp.StatusCode != 400 {
		t.Error("Expected status code 400, got ", resp.StatusCode)
		bodyString, _ := io.ReadAll(resp.Body)
		t.Error("Body: ", string(bodyString))
	}
}

func TestDeleteReq(t *testing.T) {
	app = fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})
	app.Use(recover.New())
	setRoutes(app)

	req := httptest.NewRequest("DELETE", "/dogs/1", nil)
	resp, err := app.Test(req, 10000)
	if err != nil {
		t.Error("Error sending request to Fiber: ", err)
	}

	if resp.StatusCode != 200 {
		t.Error("Expected status code 200, got ", resp.StatusCode)
		bodyString, _ := io.ReadAll(resp.Body)
		t.Error("Body: ", string(bodyString))
	}

	req = httptest.NewRequest("DELETE", "/dogs/abc", nil)
	resp, err = app.Test(req, 10000)
	if err != nil {
		t.Error("Error sending request to Fiber: ", err)
	}

	if resp.StatusCode != 400 {
		t.Error("Expected status code 400, got ", resp.StatusCode)
		bodyString, _ := io.ReadAll(resp.Body)
		t.Error("Body: ", string(bodyString))
	}
}

// Ignoreing this test for now - need proper SMTP settings in place to test which I haven't configure in GH actions yet.
// func TestInvoicePostReq(t *testing.T) {
// 	app = fiber.New(fiber.Config{
// 		Views: html.New("./views", ".html"),
// 	})
// 	app.Use(recover.New())
// 	setRoutes(app)
//
// 	req := httptest.NewRequest("POST", "/invoice/2", nil)
// 	resp, err := app.Test(req, 10000)
// 	if err != nil {
// 		t.Error("Error sending request to Fiber: ", err)
// 	}
//
// 	if resp.StatusCode != 200 {
// 		t.Error("Expected status code 200, got ", resp.StatusCode)
// 		bodyString, _ := io.ReadAll(resp.Body)
// 		t.Error("Body: ", string(bodyString))
// 	}
//
// 	req = httptest.NewRequest("POST", "/invoice/abc", nil)
// 	resp, err = app.Test(req, 10000)
// 	if err != nil {
// 		t.Error("Error sending request to Fiber: ", err)
// 	}
//
// 	if resp.StatusCode != 400 {
// 		t.Error("Expected status code 400, got ", resp.StatusCode)
// 		bodyString, _ := io.ReadAll(resp.Body)
// 		t.Error("Body: ", string(bodyString))
// 	}
// }
