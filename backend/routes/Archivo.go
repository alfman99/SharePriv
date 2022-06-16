package routes

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sharepriv/database"
	"sharepriv/entities"
	"sharepriv/middleware"
	"sharepriv/util"

	"github.com/gofiber/fiber/v2"
)

func SetArchivoRoutes(app fiber.Router) {
	// Archivo
	// Archivo Publico
	app.Get("/publico/:id/:clave", getArchivoPublico) // ACABADO
	// Middleware de autenticacion ACTIVADO
	app.Post("/publico/upload", middleware.CheckAuth, uploadArchivoPublico) // ACABADO

	// Archivo Grupo
	// Middleware de autenticacion ACTIVADO
	app.Get("/grupo/:id", middleware.CheckAuth, getArchivoGrupo) // ACABADO
	// Middleware de autenticacion ACTIVADO
	app.Post("/grupo/upload", middleware.CheckAuth, middleware.CheckGroupFormValue, uploadArchivoGrupo) // ACABADO
	// Middleware de autenticacion ACTIVADO
	app.Post("/grupo/add", middleware.CheckAuth, addArchivoGrupo) // ACABADO
}

func getArchivoPublico(c *fiber.Ctx) error {

	identificador := c.Params("id")

	claveEncriptacion := c.Params("clave")

	if len(claveEncriptacion) != 32 {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "La clave de encriptacion debe ser de 32 bytes",
		})
	}

	var archivo entities.ArchivoPublico
	if err := database.InstanciaDB.Where("id = ?", identificador).First(&archivo).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "No se encontro el archivo",
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
	archivo.Propietario = c.Locals("user").(string) // Cambiar por el usuario que subio el archivo

	if err = database.InstanciaDB.Create(&archivo).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "No se pudo crear el archivo",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Archivo creado",
		"data":    archivo,
	})
}

func getArchivoGrupo(c *fiber.Ctx) error {

	identificador := c.Params("id")

	var archivo entities.ArchivoGrupo
	if err := database.InstanciaDB.Preload("Pertenece").Where("id = ?", identificador).First(&archivo).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El archivo no existe",
		})
	}

	var usuario entities.Usuario
	if err := database.InstanciaDB.Preload("Grupos").Where("Username = ?", c.Locals("user")).Find(&usuario).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El usuario no existe",
		})
	}

	fmt.Println(usuario)

	encontrado := false

	for _, grupo := range usuario.Grupos {
		if util.ContainsGroup(archivo.Pertenece, grupo.Id) {
			encontrado = true
			break
		}
	}

	if !encontrado {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El usuario no pertenece a ningun grupo del que se puede ver este archivo",
		})
	}

	c.Context().SetContentType(archivo.Mime)

	return c.Status(200).Send(archivo.Data)
}

func uploadArchivoGrupo(c *fiber.Ctx) error {

	file, err := c.FormFile("file")
	group := c.FormValue("grupo")

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "No hay archivo en el body",
		})
	}

	if len(group) == 0 {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "No se especifico el grupo",
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

	var archivo entities.ArchivoGrupo
	archivo.Data = data
	archivo.Mime = mimeType
	archivo.Pertenece = append(archivo.Pertenece, entities.Grupo{
		Id: group,
	})
	archivo.Propietario = c.Locals("user").(string) // Cambiar por el usuario que subio el archivo

	if err = database.InstanciaDB.Create(&archivo).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "No se pudo crear el archivo",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Archivo creado",
		"data":    archivo,
	})
}

func addArchivoGrupo(c *fiber.Ctx) error {

	identificador := c.FormValue("id")
	group := c.FormValue("grupo")

	var archivo entities.ArchivoGrupo
	if err := database.InstanciaDB.Where("id = ?", identificador).First(&archivo).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El archivo no existe",
		})
	}

	if len(group) == 0 {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "No se especifico el grupo",
		})
	}

	if archivo.Propietario != c.Locals("user").(string) {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El usuario no es el propietario del archivo",
		})
	}

	archivo.Pertenece = append(archivo.Pertenece, entities.Grupo{
		Id: group,
	})

	if err := database.InstanciaDB.Save(&archivo).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "No se pudo agregar el archivo al grupo",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "El archivo se agrego al grupo",
	})

}
