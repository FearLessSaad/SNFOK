package kubernetes

import (
	"github.com/FearLessSaad/SNFOK/controllers/kubernetes/repository"
	"github.com/gofiber/fiber/v2"
)

func KubernetesInfo(router fiber.Router) {

	router.Get("/namespaces/all", func(c *fiber.Ctx) error {
		namespaces, response := repository.GetAllNamespaces()
		return c.Status(response).JSON(namespaces)
	})

}
