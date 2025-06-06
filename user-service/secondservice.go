// service-a.go
package main

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/call-get", func(c *fiber.Ctx) error {
		resp, err := http.Get("http://localhost:3001/greet")
		if err != nil {
			return c.Status(500).SendString("GET request failed: " + err.Error())
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		return c.SendString("Response from service B (GET): " + string(body))
	})

	app.Post("/call-post", func(c *fiber.Ctx) error {
		payload := []byte(`{"message":"hello from A"}`)
		resp, err := http.Post("http://localhost:3001/echo", "application/json", bytes.NewBuffer(payload))
		if err != nil {
			return c.Status(500).SendString("POST request failed: " + err.Error())
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		return c.SendString("Response from service B (POST): " + string(body))
	})

	app.Listen(":3000")
}
