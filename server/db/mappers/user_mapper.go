package db

import (
	"time"

	db "github.com/cr1m5onk1ng/nala_platform_app/db/sqlc"
)

func mapUser(user db.User) interface{} {
	return struct {
		ID               string    `json:"id"`
		Username         string    `json:"username"`
		RegistrationDate time.Time `json:"registration_date"`
		NativeLanguage   string    `json:"native_language"`
		Role             string    `json:"role"`
	}{
		ID:               user.ID,
		Username:         user.Username,
		RegistrationDate: user.RegisteredAt,
		NativeLanguage:   user.NativeLanguage,
		Role:             user.Role.String,
	}
}
