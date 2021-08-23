package domain

import _ "github.com/go-playground/validator/v10"

type MappedUser struct {
	ID       string `json:"id"`
	Username string `json:"username" validate:"required,alphanum,min=6,max=32"`
	Email    string `json:"email" validate:"required,email,min=6,max=32"`
	Password string `json:"password" validate:"required,min=10,max=32"`
	// 2 chars language code
	NativeLanguage string `json:"native_language"`
	Role           string `json:"role"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email,min=6,max=32"`
	Password string `json:"password" validate:"required,min=10,max=32"`
}
