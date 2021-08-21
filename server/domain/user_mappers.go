package domain

type MappedUser struct {
	ID               string `json:"id"`
	Username         string `json:"username"`
	Email            string `json:"email"`
	RegistrationDate string `json:"registration_date"`
	// 2 chars language code
	NativeLanguage string `json:"native_language"`
	Role           string `json:"role"`
}
