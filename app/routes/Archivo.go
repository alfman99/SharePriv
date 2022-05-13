package routes

import (
	"io/ioutil"
	"net/http"
	"sharepriv/database"
	"sharepriv/middleware"
	"sharepriv/models"
	"sharepriv/util"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func SetArchivoRoutes(app fiber.Router) {
	// Archivo
	// Archivo Publico
	app.Get("/publico/:uuid/:clave", getArchivoPublico) // ACABADO
	// Middleware de autenticacion ACTIVADO
	app.Post("/publico/upload", middleware.CheckAuth, uploadArchivoPublico) // ACABADO

	// Archivo Grupo
	// Middleware de autenticacion ACTIVADO
	app.Get("/grupo/:uuid/:clave", middleware.CheckAuth, getArchivoGrupo) // TODO
	// Middleware de autenticacion ACTIVADO
	app.Get("/grupo/upload", middleware.CheckAuth, uploadArchivoGrupo) // TODO
}

func getArchivoPublico(c *fiber.Ctx) error {

	identificador := c.Params("uuid")
	_, err := uuid.Parse(identificador)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El identificador no es un UUID",
		})
	}

	claveEncriptacion := c.Params("clave")

	if len(claveEncriptacion) != 32 {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "La clave de encriptacion debe ser de 32 bytes",
		})
	}

	var archivo models.ArchivoPublico
	database.InstanciaDB.Where("uuid = ?", identificador).First(&archivo)

	if archivo.Uuid == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El archivo no existe",
		})
	}

	decryptedFile := util.DesencriptarArchivo(archivo.Data, []byte(claveEncriptacion))

	c.Context().SetContentType(archivo.Mime)
	return c.Status(200).Send(decryptedFile)

}

func uploadArchivoPublico(c *fiber.Ctx) error {

	file, err := c.FormFile("file")

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "No hay archivo en el body",
		})
	}

	claveEncriptacion := c.FormValue("clave")

	if len(claveEncriptacion) != 32 {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "La clave de encriptacion debe ser de 32 bytes",
		})
	}

	f, err := file.Open()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "No se pudo abrir el archivo",
		})
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "No se pudo leer el archivo",
		})
	}
	mimeType := http.DetectContentType(data)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "No se pudo extraer los bytes del archivo",
		})
	}

	encryptedFile := util.EncriptarArchivo(data, []byte(claveEncriptacion))

	var archivo models.ArchivoPublico
	archivo.Data = encryptedFile
	archivo.Mime = mimeType
	archivo.PropietarioArchivo = c.Locals("user").(string) // Cambiar por el usuario que subio el archivo

	database.InstanciaDB.Create(&archivo)

	return c.JSON(&archivo.Uuid)

}

func getArchivoGrupo(c *fiber.Ctx) error {

	identificador := c.Params("uuid")
	_, err := uuid.Parse(identificador)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El identificador no es un UUID",
		})
	}

	claveEncriptacion := c.Params("clave")

	if len(claveEncriptacion) != 32 {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "La clave de encriptacion debe ser de 32 bytes",
		})
	}

	var archivo models.ArchivoGrupo
	if err := database.InstanciaDB.Where("uuid = ?", identificador).First(&archivo).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El archivo no existe",
		})
	}

	var usuario models.Usuario
	if err := database.InstanciaDB.Where("username = ?", c.Locals("username").(string)).First(&usuario); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Wtf??",
		})
	}

	grupoEncontrado := false

	for _, grupo := range usuario.Grupos {
		if grupo.Uuid == archivo.GrupoUuid {
			grupoEncontrado = true
		}
	}

	if !grupoEncontrado {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El usuario no pertenece al grupo",
		})
	}

	decryptedFile := util.DesencriptarArchivo(archivo.Data, []byte(claveEncriptacion))

	c.Context().SetContentType(archivo.Mime)
	return c.Status(200).Send(decryptedFile)

}

func uploadArchivoGrupo(c *fiber.Ctx) error {
	return nil
}
