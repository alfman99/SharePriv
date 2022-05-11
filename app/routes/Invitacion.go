package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetInvitacionRegistroRoutes(group fiber.Router) {
	// Create Invitacion Registro
	group.Post("/registro", createInvitacionRegistro) // DONE?
}

func SetInvitacionGruposRoutes(group fiber.Router) {
	// Create Invitacion Grupo
	group.Post("/grupo", createInvitacionGrupo) // DONE?
}

func createInvitacionRegistro(c *fiber.Ctx) error {
	return nil
}

func createInvitacionGrupo(c *fiber.Ctx) error {
	return nil
}
