package domain

import _ "github.com/go-playground/validator/v10"

type RefreshTokenRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email,min=6,max=32"`
	Token    string `json:"token" validate:"required"`
}
