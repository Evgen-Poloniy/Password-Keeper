package etc

type Data struct {
	Username          string `json:"username"`
	PasswordName      string `json:"password_name"`
	EncryptedPassword string `json:"encrypted_password"`
	Salt              string `json:"salt"`
	Nonce             string `json:"nonce"`
	Description       string `json:"description"`
	DateOfCreation    string `json:"date_of_creation"`
}
