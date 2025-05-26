package policies

import "github.com/gofiber/fiber/v2"

func DeployTetragonPolicy(router fiber.Router) {

	router.Post("/deploy/policy", func(c *fiber.Ctx) error {

		return c.JSON("")
	})

}
