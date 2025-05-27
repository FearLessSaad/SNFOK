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

	router.Get("/resources/all", func(c *fiber.Ctx) error {
		response, status := repository.GetAllResources()
		return c.Status(status).JSON(response)
	})

	router.Get("/get/all/labels", func(c *fiber.Ctx) error {
		response, status := repository.GetAllLabels()
		return c.Status(status).JSON(response)
	})

}
