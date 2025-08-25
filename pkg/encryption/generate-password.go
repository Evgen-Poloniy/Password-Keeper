package encryption

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func GeneratePassword(length int) (string, error) {
	const chars string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_-+={}[]|:;<>,.?/~`"

	password := make([]byte, length)
	for i := range password {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", fmt.Errorf("ошибка при генерации пароля: %v", err)
		}
		password[i] = chars[num.Int64()]
	}

	return string(password), nil
}
