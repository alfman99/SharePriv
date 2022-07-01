package main

import (
	"fmt"
	"sharepriv/database"
	"sharepriv/entities"
	"sharepriv/middleware"
	"sharepriv/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setupRoutes(app *fiber.App) {

	api := app.Group("/api")

	apiInvitaciones := api.Group("/invitaciones", middleware.CheckAuth) // /api/invitaciones protected with middleware

	// ACABADO
	apiInvitacionesRegistro := apiInvitaciones.Group("/registro")
	routes.SetInvitacionRegistroRoutes(apiInvitacionesRegistro) // /api/invitaciones/registro

	// ACABADO
	apiInvitacionesGrupos := apiInvitaciones.Group("/grupo")
	routes.SetInvitacionGruposRoutes(apiInvitacionesGrupos) // /api/invitaciones/grupos

	// ACABADO
	apiUsuarios := api.Group("/usuarios")
	routes.SetUsuarioRoutes(apiUsuarios) // /api/usuarios

	// ACABADO
	apiArchivos := api.Group("/archivos")
	routes.SetArchivoRoutes(apiArchivos) // /api/archivos

	// ACABADO
	apiGrupos := api.Group("/grupos")
	routes.SetGroupRoutes(apiGrupos) // /api/grupos

	// ACABADO
	apiAuth := api.Group("/auth")
	routes.SetAuthRoutes(apiAuth) // /api/auth

}

func prod() {

	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	database.ConnectDB()

	app := fiber.New()

	app.Use(cors.New())

	setupRoutes(app)

	app.Listen(":3000")
}

func testing() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	database.ConnectDB()

	var usuario entities.Usuario
	if err := database.InstanciaDB.Preload("PropietarioArchivoGrupo").First(&usuario).Error; err != nil {
		fmt.Println(err)
	}

	fmt.Println(usuario.PropietarioArchivoGrupo[0].Id)

}

func main() {
	// testing()
	prod()
}
