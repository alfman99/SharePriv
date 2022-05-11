package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func CheckAuth(c *fiber.Ctx) error {
	fmt.Println("CheckAuth")
	return c.Next()
}
