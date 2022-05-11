package main

import (
	"fmt"
	"sharepriv/database"
	"sharepriv/middleware"
	"sharepriv/routes"
	"sharepriv/util"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {

	api := app.Group("/api")

	apiInvitaciones := api.Group("/invitaciones", middleware.CheckAuth) // /api/invitaciones protected with middleware
	apiInvitacionesRegistro := apiInvitaciones.Group("/registro")
	routes.SetInvitacionRegistroRoutes(apiInvitacionesRegistro) // /api/invitaciones/registro

	apiInvitacionesGrupos := apiInvitaciones.Group("/grupos")
	routes.SetInvitacionGruposRoutes(apiInvitacionesGrupos) // /api/invitaciones/grupos

	apiUsuarios := api.Group("/usuarios")
	routes.SetUsuarioRoutes(apiUsuarios) // /api/usuarios

	apiArchivos := api.Group("/archivos")
	routes.SetArchivoRoutes(apiArchivos) // /api/archivos

	apiGrupos := api.Group("/grupos")
	routes.SetGroupRoutes(apiGrupos) // /api/grupos

}

func prod() {
	database.ConnectDB()

	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}

func testing() {
	test := util.EncriptarArchivo([]byte("hola"), []byte("passphrasewhichneedstobe32bytes!"))

	fmt.Println(string(test))

	fmt.Println(string(util.DesencriptarArchivo(test, []byte("passphrasewhichneedstobe32bytes!"))))
}

func main() {
	// testing()
	prod()
}
