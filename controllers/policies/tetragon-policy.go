package policies

import (
	"fmt"

	"github.com/FearLessSaad/SNFOK/controllers/policies/repository"
	"github.com/gofiber/fiber/v2"
)

func DeployTetragonPolicy(router fiber.Router) {
	router.Get("/deploy/:id/:namespace/:label/", func(c *fiber.Ctx) error {
		id := c.AllParams()["id"]
		namespace := c.AllParams()["namespace"]
		label := c.AllParams()["label"]
		if id == "" || namespace == "" || label == "" {
			return c.Status(fiber.StatusBadRequest).JSON("")
		}
		res, status := repository.DeployPolicy(namespace, label, id)
		fmt.Println(res)
		return c.Status(status).JSON(res)
	})

	router.Get("/get/all", func(c *fiber.Ctx) error {
		namespaces, response := repository.GetAllImplimentedPolicies()
		return c.Status(response).JSON(namespaces)
	})

	router.Get("/delete/:id", func(c *fiber.Ctx) error {
		namespaces, response := repository.DeletePolicy(c.AllParams()["id"])
		return c.Status(response).JSON(namespaces)
	})

}
