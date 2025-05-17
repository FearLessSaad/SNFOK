package auth

import (
	"time"

	"github.com/FearLessSaad/SNFOK/constants/message"
	"github.com/FearLessSaad/SNFOK/constants/response"
	"github.com/FearLessSaad/SNFOK/controllers/auth/dto"
	"github.com/FearLessSaad/SNFOK/controllers/auth/repository"
	"github.com/FearLessSaad/SNFOK/tooling/global_dto"
	"github.com/FearLessSaad/SNFOK/tooling/security/token"
	"github.com/FearLessSaad/SNFOK/tooling/security/validation"
	"github.com/gofiber/fiber/v2"
)

func AuthController(router fiber.Router) {

	router.Post("/login", func(c *fiber.Ctx) error {
		details := new(dto.LoginDetails)
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

		data, status := repository.Login(*details)

		if status != fiber.StatusOK {
			return c.Status(status).JSON(data)
		}

		c.Cookie(&fiber.Cookie{
			Name:     "access_token",
			Value:    data.Data.AccessToken,
			HTTPOnly: true,
			Secure:   false,
			SameSite: "Strict",
			Expires:  time.Now().Add(2 * 60 * time.Minute),
		})

		data.Data = nil
		return c.Status(status).JSON(data)
	})

	router.Get("/validate", func(c *fiber.Ctx) error {
		access_cookie := c.Cookies("access_token")
		if access_cookie == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized access",
			})
		}

		valid, errMsg, _ := token.VerifyAccessTokenAndGetClaims(access_cookie)
		if !valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": errMsg,
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{})
	})

}
