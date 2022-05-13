package routes

import (
	"sharepriv/database"
	"sharepriv/entities"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SetUsuarioRoutes(app fiber.Router) {
	// Create usuario
	app.Post("/", createUser)

	// Get usuario by username
	app.Get("/:username", getUser) // DONE!
}

func createUser(c *fiber.Ctx) error {

	payload := struct {
		Username   string `json:"username"`
		Password   string `json:"password"`
		Invitacion string `json:"invitacion"`
	}{}

	if err := c.BodyParser(&payload); err != nil || payload.Username == "" || payload.Password == "" || payload.Invitacion == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El body no tiene el formato correcto",
		})
	}

	var invitacion entities.InvitacionRegistro

	if err := database.InstanciaDB.Where("codigo = ?", payload.Invitacion).First(&invitacion).Error; err != nil {
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

	var usuario entities.Usuario

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

	if username == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El username no puede estar vacio",
		})
	}

	var usuario entities.Usuario

	if err := database.InstanciaDB.Where("username = ?", username).First(&usuario).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Usuario no encontrado",
		})
	}

	return c.Status(200).JSON(usuario)

}
