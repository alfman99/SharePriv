package routes

import (
	"io/ioutil"
	"sharepriv/database"
	"sharepriv/models"
	"sharepriv/util"

	"github.com/gofiber/fiber/v2"
)

func SetArchivoRoutes(app fiber.Router) {
	// Archivo
	// Archivo Publico

	app.Get("/publico/:uuid", getArchivoPublico) // TODO
	// Archivo Grupo
	app.Get("/grupo/:uuid", getArchivoGrupo) // TODO

	app.Post("/publico/upload", uploadArchivoPublico) // TODO
}

// TODO: Get archivo desencriptado
func getArchivoPublico(c *fiber.Ctx) error {
	return nil
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

	fileType, err := util.GetFileContentType(f)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "No se pudo determinar el tipo de archivo",
		})
	}

	if fileType != "image/png" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El archivo no es una imagen",
		})
	}

	byteContainer, err := ioutil.ReadAll(f)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "No se pudo extraer los bytes del archivo",
		})
	}

	encryptedFile := util.EncriptarArchivo(byteContainer, []byte(claveEncriptacion))

	var clave models.Clave
	clave.Clave = claveEncriptacion
	database.InstanciaDB.Create(&clave)

	var archivo models.ArchivoPublico
	archivo.Data = encryptedFile
	archivo.ClaveClave = claveEncriptacion
	archivo.PropietarioArchivo = "test" // Cambiar por el usuario que subio el archivo

	database.InstanciaDB.Create(&archivo)

	return c.JSON(&archivo.Uuid)

}
