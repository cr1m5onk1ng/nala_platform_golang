package domain

import _ "github.com/go-playground/validator/v10"

type MappedResource struct {
	Url string `json:"url" validate:"required,url"`
	// 2 chars language code
	Language  string `json:"language" validate:"required,min=2,max=2,alpha"`
	MediaType string `json:"media_type" validate:"required,alpha"`
	Category  string `json:"category" validate:"required,alpha"`
}
