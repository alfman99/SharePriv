package main

import (
	"sharepriv/database"

	"github.com/gofiber/fiber/v2"
)

func test(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}

func main() {
	database.ConnectDB()

	app := fiber.New()

	app.Get("/api", test)

	app.Listen(":3000")

}
