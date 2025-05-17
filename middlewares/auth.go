package middlewares

import (
	"github.com/FearLessSaad/SNFOK/tooling/security/token"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	access_token := c.Cookies("access_token")
	if access_token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized",
		})
	}
	isValid, err, claims := token.VerifyAccessTokenAndGetClaims(access_token)
	if err != "" || !isValid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized",
		})
	}

	c.Locals("user_id", claims.UserID)

	return c.Next()
}
