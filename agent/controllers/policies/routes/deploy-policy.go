package routes

import (
	"github.com/FearLessSaad/SNFOK/agent/controllers/policies/features"
	"github.com/FearLessSaad/SNFOK/constants/message"
	"github.com/FearLessSaad/SNFOK/constants/response"
	"github.com/FearLessSaad/SNFOK/shared/agent_dto"
	"github.com/FearLessSaad/SNFOK/tooling/global_dto"
	"github.com/gofiber/fiber/v2"
)

type PolicyPathRequest struct {
	Path string `json:"path"`
}

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

	router.Post("/delete", func(c *fiber.Ctx) error {

		details := new(PolicyPathRequest)
		if err := c.BodyParser(details); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(global_dto.Response[string]{
				Status:  "error",
				Message: message.INVALID_REQUEST_PAYLOAD,
				Data:    nil,
				Meta: &global_dto.Meta{
					Code: response.INVALID_REQUEST_PAYLOAD,
				},
			})
		}
		_, err := features.DeletePolicy(details.Path)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return c.Status(fiber.StatusOK).JSON("")
	})
}
