package routes

import (
	"sharepriv/database"
	"sharepriv/entities"
	"sharepriv/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetGroupRoutes(app fiber.Router) {
	// Get group by id
	app.Get("/:id", middleware.CheckAuth, getGroup) // ACABADO

	// Create group
	app.Post("/", middleware.CheckAuth, createGroup) // ACABADO

	// Unirse a grupo
	app.Post("/join", middleware.CheckAuth, joinGroup) // ACABADO
}

func getGroup(c *fiber.Ctx) error {

	identifier := c.Params("id")

	var grupo entities.Grupo
	if err := database.InstanciaDB.Preload("Usuarios").Preload("Archivos").Where("id = ?", identifier).First(&grupo).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El grupo no existe",
		})
	}

	participante := false

	for _, usuario := range grupo.Usuarios {
		if usuario.Username == c.Locals("user").(string) {
			participante = true
			break
		}
	}

	if !participante {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El usuario no pertenece al grupo",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   grupo,
	})
}

func createGroup(c *fiber.Ctx) error {

	nombre := c.FormValue("nombre")

	if len(nombre) == 0 {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El nombre del grupo no puede estar vacio",
		})
	}

	var usuario entities.Usuario
	if err := database.InstanciaDB.Where("username = ?", c.Locals("user").(string)).First(&usuario).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El usuario no existe",
		})
	}

	grupo := entities.Grupo{
		Nombre:      nombre,
		Propietario: c.Locals("user").(string),
		Usuarios: []entities.Usuario{
			{
				Username: c.Locals("user").(string),
			},
		},
	}

	if err := database.InstanciaDB.Create(&grupo).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Error al crear el grupo",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "Grupo creado",
		"data": fiber.Map{
			"id":     grupo.Id,
			"nombre": grupo.Nombre,
		},
	})
}

func joinGroup(c *fiber.Ctx) error {

	invitacionCode := c.FormValue("invitacion")

	if len(invitacionCode) == 0 {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "La invitacion no puede estar vacia",
		})
	}

	var invitacion entities.InvitacionGrupo
	if err := database.InstanciaDB.Where("codigo = ?", invitacionCode).First(&invitacion).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "La invitacion no existe",
		})
	}

	var grupo entities.Grupo
	if err := database.InstanciaDB.Preload("Usuarios").Preload("Archivos").Where("id = ?", invitacion.GrupoId).First(&grupo).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El grupo no existe",
		})
	}

	participante := false

	for _, usuario := range grupo.Usuarios {
		if usuario.Username == c.Locals("user").(string) {
			participante = true
			break
		}
	}

	if participante {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El usuario ya pertenece al grupo",
		})
	}

	grupo.Usuarios = append(grupo.Usuarios, entities.Usuario{
		Username: c.Locals("user").(string),
	})

	if err := database.InstanciaDB.Save(&grupo).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Error al unirse al grupo",
		})
	}

	invitacion.Usos = invitacion.Usos + 1

	if err := database.InstanciaDB.Save(&invitacion).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Error al actualizar la invitacion",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Usuario unido al grupo",
	})
}
