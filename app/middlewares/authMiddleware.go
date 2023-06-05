package middlewares

import (
	"go-folder-sample/app/helpers"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Authentication() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorizationHeader := c.Get("Authorization")
		if authorizationHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "No Authorization header provided"})
		}

		// Split the Authorization header to extract the token
		authHeaderParts := strings.Split(authorizationHeader, " ")
		if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid Authorization header format"})
		}

		clientToken := authHeaderParts[1]

		claims, err := helpers.ValidateAccessToken(clientToken)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
		}
		// Modify the request if needed
		// Example: Setting a custom request header
		c.Request().Header.Set("identifier", claims.Id)

		return c.Next()
	}
}
