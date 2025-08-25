package etc

type UserSettings struct {
	CurrentUsername         string `json:"current_username"`
	CurrentUserPasswordHash string `json:"current_password_hash"`
	PasswordGenerationLenth int    `json:"password_generation_lenth"`
}

var Settings UserSettings
