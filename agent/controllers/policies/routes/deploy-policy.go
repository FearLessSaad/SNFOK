package routes

import (
	"github.com/FearLessSaad/SNFOK/agent/controllers/policies/features"
	"github.com/FearLessSaad/SNFOK/shared/agent_dto"
	"github.com/gofiber/fiber/v2"
)

func DeployPolicy(router fiber.Router) {

	router.Post("/deplye/policy", func(c *fiber.Ctx) error {
		details := new(agent_dto.DeployPolicy)
		if err := c.BodyParser(details); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("")
		}

		path, err := features.DeployPolicy(details.FilePath, details.Namespace, details.AppLabel)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"policy_path": path,
		})
	})
}
