package kubernetes

import (
	"github.com/gofiber/fiber/v2"
)

func KubernetesInfo(router fiber.Router) {

	router.Get("/namespaces/all", func(c *fiber.Ctx) error {

		return c.JSON("")
	})

}
