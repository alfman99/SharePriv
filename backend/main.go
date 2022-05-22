package main

import (
	"fmt"
	"sharepriv/database"
	"sharepriv/middleware"
	"sharepriv/routes"
	"sharepriv/util"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
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

	setupRoutes(app)

	app.Listen(":3000")
}

func testing() {
	/*test := util.EncriptarArchivo([]byte("♀♀µ┼"), []byte("passphrasewhichneedstobe32bytes!"))

	fmt.Println(string(test))

	fmt.Println(util.DesencriptarArchivo(test, []byte("passphrasewhichneedstobe32bytes!")))*/

	fmt.Println(util.GenerateRandomString(16))
	fmt.Println(util.GenerateRandomString(16))

}

func main() {
	// testing()
	prod()
}
