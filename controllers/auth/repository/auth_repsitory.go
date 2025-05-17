package repository

import (
	"github.com/FearLessSaad/SNFOK/constants/response"
	"github.com/FearLessSaad/SNFOK/controllers/auth/dto"
	"github.com/FearLessSaad/SNFOK/controllers/auth/persistance"
	"github.com/FearLessSaad/SNFOK/tooling/global_dto"
	"github.com/FearLessSaad/SNFOK/tooling/security/token"
	"github.com/gofiber/fiber/v2"
)

func Login(data dto.LoginDetails) (global_dto.Response[dto.JWTToken], int) {

	status, user := persistance.GetUserByEmailAddress(data.Email)

	if !status {
		return global_dto.Response[dto.JWTToken]{
			Status:  "error",
			Message: "Invalid email address or token.",
			Data:    nil,
			Meta: &global_dto.Meta{
				Code: response.INVALID_EMAIL_PASSWORD,
			},
		}, fiber.StatusUnauthorized
	}

	if data.Token != user.Token {
		return global_dto.Response[dto.JWTToken]{
			Status:  "error",
			Message: "Invalid email address or token.",
			Data:    nil,
			Meta: &global_dto.Meta{
				Code: response.INVALID_EMAIL_PASSWORD,
			},
		}, fiber.StatusUnauthorized
	}

	token, err := token.GenerateJWTTokens(user.ID, user.Designation)

	if err != "" {
		return global_dto.Response[dto.JWTToken]{
			Status:  "error",
			Message: "Invalid email address or token.",
			Data:    nil,
			Meta: &global_dto.Meta{
				Code: response.EXECUTION_ERROR,
			},
		}, fiber.StatusUnauthorized
	}

	return global_dto.Response[dto.JWTToken]{
		Status:  "success",
		Message: "Logged in successfully!",
		Data:    &token,
		Meta: &global_dto.Meta{
			Code: response.LOGIN_SUCCESS,
		},
	}, fiber.StatusOK
}
