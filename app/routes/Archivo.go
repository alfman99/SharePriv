package routes

import (
	"io/ioutil"
	"net/http"
	"sharepriv/database"
	"sharepriv/entities"
	"sharepriv/middleware"
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
	app.Get("/grupo/:uuid/:clave", middleware.CheckAuth, getArchivoGrupo) // ACABADO
	// Middleware de autenticacion ACTIVADO
	app.Post("/grupo/upload", middleware.CheckAuth, middleware.CheckGroupFormValue, uploadArchivoGrupo) // ACABADO
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

	var archivo entities.ArchivoPublico
	database.InstanciaDB.Where("uuid = ?", identificador).First(&archivo)

	if archivo.Uuid == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El archivo no existe",
		})
	}

	decryptedFile, err := util.DesencriptarArchivo(archivo.Data, []byte(claveEncriptacion))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "No se pudo desencriptar el archivo, la clave es incorrecta",
		})
	}

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

	var archivo entities.ArchivoPublico
	archivo.Data = encryptedFile
	archivo.Mime = mimeType
	archivo.PropietarioArchivo = c.Locals("user").(string) // Cambiar por el usuario que subio el archivo

	if err = database.InstanciaDB.Create(&archivo).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "No se pudo crear el archivo",
		})
	}

	return c.JSON(&archivo.Uuid)
}

func getArchivoGrupo(c *fiber.Ctx) error {

	identificador := c.Params("uuid")

	_, err := uuid.Parse(identificador)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El identificador del archivo no es un UUID",
		})
	}

	claveEncriptacion := c.Params("clave")

	if len(claveEncriptacion) != 32 {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "La clave de encriptacion debe ser de 32 bytes",
		})
	}

	var archivo entities.ArchivoGrupo
	if err := database.InstanciaDB.Where("uuid = ?", identificador).First(&archivo).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El archivo no existe",
		})
	}

	var grupo entities.Grupo
	if err := database.InstanciaDB.Preload("Usuarios").Where("uuid = ?", archivo.GrupoUuid).First(&grupo).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El grupo del archivo no existe",
		})
	}

	encontrado := false

	for _, usuario := range grupo.Usuarios {
		if usuario.Username == c.Locals("user").(string) {
			encontrado = true
			break
		}
	}

	if !encontrado {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El usuario no pertenece al grupo",
		})
	}

	decryptedFile, err := util.DesencriptarArchivo(archivo.Data, []byte(claveEncriptacion))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "No se pudo desencriptar el archivo, la clave es incorrecta",
		})
	}

	c.Context().SetContentType(archivo.Mime)

	return c.Status(200).Send(decryptedFile)
}

func uploadArchivoGrupo(c *fiber.Ctx) error {

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

	var archivo entities.ArchivoGrupo
	archivo.Data = encryptedFile
	archivo.Mime = mimeType
	archivo.GrupoUuid = c.FormValue("grupo")
	archivo.PropietarioArchivo = c.Locals("user").(string) // Cambiar por el usuario que subio el archivo

	if err = database.InstanciaDB.Create(&archivo).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "No se pudo crear el archivo",
		})
	}

	return c.JSON(&archivo.Uuid)
}
