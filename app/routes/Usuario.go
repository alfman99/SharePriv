package routes

import (
	"sharepriv/database"
	"sharepriv/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SetUsuarioRoutes(app fiber.Router) {
	// Get usuario by username
	app.Get("/:username", getUser) // DONE!

	// Create usuario
	app.Post("/", createUser)
}

func createUser(c *fiber.Ctx) error {

	payload := struct {
		Username   string `json:"username"`
		Password   string `json:"password"`
		Invitacion string `json:"invitacion"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El body no tiene el formato correcto",
		})
	}

	var invitacion models.InvitacionRegistro
	invitacion.Codigo = payload.Invitacion

	if err := database.InstanciaDB.First(&invitacion).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El codigo de invitacion no es valido",
		})
	}

	if invitacion.Usos >= invitacion.MaximoUsos {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El codigo de invitacion ya ha sido usado el maximo de veces",
		})
	}

	if invitacion.FechaCaducidad.Before(time.Now()) {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El codigo de invitacion ha expirado",
		})
	}

	var usuario models.Usuario

	usuario.Username = payload.Username
	usuario.Password = payload.Password
	usuario.InvitacionRegistroCodigo = payload.Invitacion

	if err := database.InstanciaDB.Create(&usuario).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El usuario ya existe",
		})
	}

	invitacion.Usos++

	if err := database.InstanciaDB.Save(&invitacion).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Error al actualizar el codigo de invitacion",
		})
	}

	return c.Status(200).JSON(usuario)

}

func getUser(c *fiber.Ctx) error {
	username := c.AllParams()["username"]

	var usuario models.Usuario

	if err := database.InstanciaDB.Where("username = ?", username).First(&usuario).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Usuario no encontrado",
		})
	}

	return c.Status(200).JSON(usuario)

}
