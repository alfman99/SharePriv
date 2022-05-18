package middleware

import (
	"sharepriv/database"
	"sharepriv/entities"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

var JwtKey = []byte("fas8df8as3ll")

func CheckAuth(c *fiber.Ctx) error {

	token := c.Cookies("token")

	if token == "" {
		return c.Status(401).JSON(fiber.Map{
			"status":  "error",
			"message": "No esta autenticado",
		})
	}

	claims := jwt.MapClaims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtKey), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return c.Status(401).JSON(fiber.Map{
				"status":  "error",
				"message": "Error firma no valida del token",
			})
		}
	}

	if !tkn.Valid {
		return c.Status(401).JSON(fiber.Map{
			"status":  "error",
			"message": "Token invalido",
		})
	}

	if claims, ok := tkn.Claims.(jwt.MapClaims); ok && tkn.Valid {
		c.Locals("user", claims["user"].(string))
	} else {
		return c.Status(401).JSON(fiber.Map{
			"status":  "error",
			"message": "Token has no user?",
		})
	}

	return c.Next()
}

func CheckGroupFormValue(c *fiber.Ctx) error {

	grupo := c.FormValue("grupo")

	/*_, err := uuid.Parse(grupo)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El identificador del grupo no es un UUID",
		})
	}*/

	username := c.Locals("user").(string)

	var user entities.Usuario
	if err := database.InstanciaDB.Preload("Grupos").Where("username = ?", username).First(&user).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El usuario no existe",
		})
	}

	grupoEncontrado := false

	for _, grp := range user.Grupos {
		if grp.Id == grupo {
			grupoEncontrado = true
			break
		}
	}

	if !grupoEncontrado {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "El usuario no pertenece al grupo",
		})
	}

	return c.Next()
}
