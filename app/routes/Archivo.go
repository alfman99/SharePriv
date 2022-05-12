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
	app.Get("/publico/:uuid", getArchivoPublico) // TODO

	// Archivo Grupo
	// Middleware de autenticacion ACTIVADO
	app.Get("/grupo/:uuid", middleware.CheckAuth, getArchivoGrupo) // TODO

	// Middleware de autenticacion ACTIVADO
	app.Post("/publico/upload", middleware.CheckAuth, uploadArchivoPublico) // TODO
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

	claveEncriptacion := []byte(c.FormValue("clave"))

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

	decryptedFile := util.DesencriptarArchivo(archivo.Data, claveEncriptacion)

	c.Context().SetContentType(archivo.Mime)
	return c.Status(200).Send(decryptedFile)

}

// TODO: Get archivo de un grupo desencriptado
func getArchivoGrupo(c *fiber.Ctx) error {
	return nil
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
	archivo.PropietarioArchivo = "test" // Cambiar por el usuario que subio el archivo

	database.InstanciaDB.Create(&archivo)

	return c.JSON(&archivo.Uuid)

}
