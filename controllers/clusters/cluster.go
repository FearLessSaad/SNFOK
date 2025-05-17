package clusters

import (
	"github.com/FearLessSaad/SNFOK/constants/message"
	"github.com/FearLessSaad/SNFOK/constants/response"
	"github.com/FearLessSaad/SNFOK/controllers/clusters/dto"
	"github.com/FearLessSaad/SNFOK/controllers/clusters/repository"
	"github.com/FearLessSaad/SNFOK/tooling/global_dto"
	"github.com/FearLessSaad/SNFOK/tooling/security/validation"
	"github.com/gofiber/fiber/v2"
)

func ClusterInfo(router fiber.Router) {

	router.Get("/all", func(c *fiber.Ctx) error {
		response, status := repository.GetAllClusters()
		return c.Status(status).JSON(response)
	})

	router.Post("/create", func(c *fiber.Ctx) error {
		details := new(dto.ClusterRequest)
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
		if errs := validation.ValidateStruct(details); len(errs) > 0 {
			errors := make([]any, len(errs))
			for i, err := range errs {
				errors[i] = err
			}
			return c.Status(fiber.StatusUnprocessableEntity).JSON(global_dto.Response[string]{
				Status:  "error",
				Message: message.FAILED_DATA_VALIDATION,
				Errors:  errors,
				Data:    nil,
				Meta: &global_dto.Meta{
					Code: response.FAILED_DATA_VALIDATION,
				},
			})
		}

		user_id := c.Locals("user_id").(string)
		response, status := repository.AddNewCluster(*details, user_id)
		return c.Status(status).JSON(response)
	})
}
