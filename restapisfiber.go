package main

import (
	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{ID: 1, Name: "Alice", Email: "alice@example.com"},
	{ID: 2, Name: "Bob", Email: "bob@example.com"},
}

func getApi() {

}
func postApi() {

}
func main() {

	app := fiber.New()

	app.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(users)
	})

	app.Post("/addUser", func(c *fiber.Ctx) error {
		newUser := new(User)
		if err := c.BodyParser(newUser); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "cannot parse JSON",
			})
		}
		newUser.ID = len(users) + 1
		users = append(users, *newUser)
		return c.Status(fiber.StatusCreated).JSON(newUser)
	})
	app.Listen(":3000")
}
