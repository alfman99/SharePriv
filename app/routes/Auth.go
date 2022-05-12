package routes

import (
	"sharepriv/database"
	"sharepriv/middleware"
	"sharepriv/models"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gofiber/fiber/v2"
)

func SetAuthRoutes(app fiber.Router) {
	// Login user
	app.Post("/login", setLogin) // TODO
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

	var user models.Usuario
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

	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Token generado",
	})

}
