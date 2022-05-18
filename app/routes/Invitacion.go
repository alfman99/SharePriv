package routes

import (
	"sharepriv/database"
	"sharepriv/entities"
	"sharepriv/util"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SetInvitacionRegistroRoutes(group fiber.Router) {
	// Create Invitacion Registro
	group.Post("/crear", createInvitacionRegistro) // ACABADO
}

func SetInvitacionGruposRoutes(group fiber.Router) {
	// Create Invitacion Grupo
	group.Post("/crear", createInvitacionGrupo) // ACABADO
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

	var invitacion entities.InvitacionRegistro

	invitacion.Codigo = util.GenerateRandomString(16)
	invitacion.FechaCaducidad = fechaVal
	invitacion.MaximoUsos = uint(maximoUsos)
	invitacion.Propietario = c.Locals("user").(string)

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

	payload := struct {
		FechaCaducidad string `json:"fecha_caducidad"`
		MaximoUsos     string `json:"maximo_usos"`
		GrupoUuid      string `json:"grupo_uuid"`
	}{}

	if err := c.BodyParser(&payload); err != nil || payload.FechaCaducidad == "" || payload.MaximoUsos == "" || payload.GrupoUuid == "" {
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

	// Check if string is valid uuid
	/*if _, err := uuid.Parse(payload.GrupoUuid); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El formato del uuid del grupo no es valido",
		})
	}*/

	// Check if group exists
	var grupo entities.Grupo
	if err := database.InstanciaDB.Where("id = ?", payload.GrupoUuid).First(&grupo).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El grupo para el que quieres crear la invitación no existe",
		})
	}

	if grupo.PropietarioUsername != c.Locals("user").(string) {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "No eres el propietario del grupo",
		})
	}

	var invitacion entities.InvitacionGrupo

	invitacion.Codigo = util.GenerateRandomString(16)
	invitacion.FechaCaducidad = fechaVal
	invitacion.MaximoUsos = uint(maximoUsos)
	invitacion.Propietario = c.Locals("user").(string)
	invitacion.GrupoId = payload.GrupoUuid

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
