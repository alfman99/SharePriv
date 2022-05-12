package routes

import (
	"sharepriv/database"
	"sharepriv/models"
	"sharepriv/util"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SetInvitacionRegistroRoutes(group fiber.Router) {
	// Create Invitacion Registro
	group.Post("/crear", createInvitacionRegistro) // DONE?
}

func SetInvitacionGruposRoutes(group fiber.Router) {
	// Create Invitacion Grupo
	group.Post("/crear", createInvitacionGrupo) // DONE?
}

func createInvitacionRegistro(c *fiber.Ctx) error {

	payload := struct {
		FechaCaducidad string `json:"fecha_caducidad"`
		MaximoUsos     string `json:"maximo_usos"`
	}{}

	if err := c.BodyParser(&payload); err != nil || payload.FechaCaducidad == "" || payload.MaximoUsos == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El body no tiene el formato correcto",
		})
	}

	fechaVal, err := time.Parse("2006-01-02", payload.FechaCaducidad)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El formato de la fecha de caducidad no es valido",
		})
	}

	maximoUsos, err := strconv.Atoi(payload.MaximoUsos)

	if err != nil || maximoUsos <= 1 {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El formato del maximo de usos no es valido",
		})
	}

	var invitacion models.InvitacionRegistro

	invitacion.Codigo = util.GenerateRandomString(16)
	invitacion.FechaCaducidad = fechaVal
	invitacion.MaximoUsos = uint(maximoUsos)
	invitacion.Propietario = c.Get("user")

	if err := database.InstanciaDB.Create(&invitacion).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Error al crear la invitacion",
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   invitacion,
	})

}

func createInvitacionGrupo(c *fiber.Ctx) error {
	return nil
}
