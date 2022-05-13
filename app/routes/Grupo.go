package routes

import (
	"sharepriv/database"
	"sharepriv/entities"
	"sharepriv/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func SetGroupRoutes(app fiber.Router) {
	// Get group by id
	app.Get("/:uuid", middleware.CheckAuth, getGroup) // ACABADO

	// Create group
	app.Post("/", middleware.CheckAuth, createGroup) // ACABADO
}

// TODO: Get user group info
func getGroup(c *fiber.Ctx) error {

	identifier := c.Params("uuid")

	_, err := uuid.Parse(identifier)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El identificador no es un UUID",
		})
	}

	var grupo entities.Grupo
	if err := database.InstanciaDB.Preload("Usuarios").Preload("Archivos").Where("uuid = ?", identifier).First(&grupo).Error; err != nil {
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
		Nombre:              nombre,
		PropietarioUsername: c.Locals("user").(string),
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
			"uuid":   grupo.Uuid,
			"nombre": grupo.Nombre,
		},
	})

}
