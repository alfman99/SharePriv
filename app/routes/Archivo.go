package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetArchivoRoutes(app fiber.Router) {
	// Archivo
	// Archivo Publico

	app.Get("/publico/:uuid", getArchivoPublico) // TODO
	// Archivo Grupo
	app.Get("/grupo/:uuid", getArchivoGrupo) // TODO
}

// TODO: Get archivo desencriptado
func getArchivoPublico(c *fiber.Ctx) error {
	return nil
}

// TODO: Get archivo de un grupo desencriptado
func getArchivoGrupo(c *fiber.Ctx) error {
	return nil
}
