package routes

import (
	"fmt"
	"sharepriv/database"
	"sharepriv/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

	payload := struct {
		Invitacion     string `json:"invitacion"`
		FechaCaducidad string `json:"fechaCaducidad"`
		MaximoUsos     uint   `json:"maximoUsos"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		fmt.Println(err)
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El body no tiene el formato correcto",
		})
	}

	var invitacion models.InvitacionRegistro
	invitacion.Codigo = payload.Invitacion
	valFecha, err := time.Parse("01-02-2006", payload.FechaCaducidad)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El formato de la fecha de caducidad no es valido",
		})
	}

	invitacion.FechaCaducidad = valFecha
	invitacion.MaximoUsos = payload.MaximoUsos

	if err := database.InstanciaDB.Create(&invitacion).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El codigo de invitacion ya existe",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Invitacion creada",
	})

}

func createInvitacionGrupo(c *fiber.Ctx) error {

	payload := struct {
		Invitacion     string `json:"invitacion"`
		FechaCaducidad string `json:"fechaCaducidad"`
		MaximoUsos     uint   `json:"maximoUsos"`
		Grupo          string `json:"grupo"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		fmt.Println(err)
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El body no tiene el formato correcto",
		})
	}

	var invitacion models.InvitacionGrupo
	invitacion.Codigo = payload.Invitacion
	valFecha, err := time.Parse("01-02-2006", payload.FechaCaducidad)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El formato de la fecha de caducidad no es valido",
		})
	}

	invitacion.FechaCaducidad = valFecha
	invitacion.MaximoUsos = payload.MaximoUsos

	_, err = uuid.Parse(payload.Grupo)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El uuid del grupo no tiene un formato valido",
		})
	}

	var grupoCreacionInv models.Grupo
	if err := database.InstanciaDB.Where("uuid = ?", payload.Grupo).First(&grupoCreacionInv).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El grupo para el que quieres crear la invitación no existe",
		})
	}

	invitacion.GrupoUuid = payload.Grupo

	if err := database.InstanciaDB.Create(&invitacion).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El codigo de invitacion ya existe",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Invitacion creada",
	})

}
