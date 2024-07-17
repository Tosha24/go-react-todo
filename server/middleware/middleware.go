package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tosha24/todo/utils"
)


func JWTMiddleware(c *fiber.Ctx) error {
	// Get the token from the authorization header
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "No token found",
		})
	}

	// Authenticate the token
	userId, err := utils.AuthenticateJWTToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	// Set the userId to the local context
	c.Locals("userId", userId)

	return c.Next()
}