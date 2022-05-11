package routes

import "github.com/gofiber/fiber/v2"

func SetGroupRoutes(app fiber.Router) {
	// Get group by id
	app.Get("/:id", getGroup) // TODO

	// Create group
	app.Post("/", createGroup) // TODO
}

// TODO: Get user group info
func getGroup(c *fiber.Ctx) error {
	return nil
}

// TODO: Create group
func createGroup(c *fiber.Ctx) error {
	return nil
}
