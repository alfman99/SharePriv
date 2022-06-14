package routes

import (
	"fmt"
	"sharepriv/database"
	"sharepriv/entities"
	"sharepriv/middleware"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/gofiber/fiber/v2"
)

func SetAuthRoutes(app fiber.Router) {
	// Login user
	app.Post("/login", setLogin)                            // ACABADO
	app.Get("/validate", middleware.CheckAuth, getValidate) // ACABADO

	app.Get("/grupos", middleware.CheckAuth, getUserGrupos) // ACABADO
}

func setLogin(c *fiber.Ctx) error {

	payload := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := c.BodyParser(&payload); err != nil || payload.Username == "" || payload.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El body no tiene el formato correcto",
		})
	}

	var user entities.Usuario
	if err := database.InstanciaDB.Where("username = ?", payload.Username).First(&user).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El usuario no existe",
		})
	}

	if user.Password != payload.Password {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "La contraseña es incorrecta",
		})
	}

	expirationTime := time.Now().Add(time.Hour * 24 * 7)

	claims := &jwt.MapClaims{
		"user": payload.Username,
		"StandardClaims": jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(middleware.JwtKey)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Error al generar el token",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Token generado",
		"token":   tokenString,
	})
}

func getValidate(c *fiber.Ctx) error {

	tokenString := c.Get("Authorization")

	if tokenString == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "No se ha enviado el token",
		})
	}

	claims := jwt.MapClaims{}

	tkn, _ := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(middleware.JwtKey), nil
	})

	if claims, ok := tkn.Claims.(jwt.MapClaims); ok && tkn.Valid {
		return c.Status(200).JSON(fiber.Map{
			"status": "success",
			"datos":  claims,
		})
	}

	return c.Status(401).JSON(fiber.Map{
		"status":  "error",
		"message": "Token invalido",
	})

}

func getUserGrupos(c *fiber.Ctx) error {

	var user entities.Usuario
	if err := database.InstanciaDB.Preload("Grupos").Where("username = ?", c.Locals("user")).First(&user).Error; err != nil {

		fmt.Println(err)

		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El usuario no existe",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   user.Grupos,
	})

}
