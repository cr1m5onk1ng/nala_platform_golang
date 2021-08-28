package domain

import _ "github.com/go-playground/validator/v10"

type MappedUserPost struct {
	ID              string `json:"id"`
	UserID          string `json:"user_id" validate:"required"`
	ResourceID      int64  `json:"resource_id" validate:"required"`
	PostTitle       string `json:"post_title" validate:"required,alphanum"`
	PostDescription string `json:"post_description" validate:"alphanum"`
}
