package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/greet", func(c *fiber.Ctx) error {

		return c.JSON("Hello from Service B!")

	})

	app.Post("/add", func(c *fiber.Ctx) error {

		body := c.Body()
		if len(body) == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Request bdy is Empty",
			})
		}
		return c.JSON(body)
	})

	app.Listen(":3001")
}
