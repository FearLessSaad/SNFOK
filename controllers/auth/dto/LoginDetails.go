package dto

type LoginDetails struct {
	Email string `json:"email" validate:"required,email"`
	Token string `json:"token" validate:"required,min=64,max=64"`
}
