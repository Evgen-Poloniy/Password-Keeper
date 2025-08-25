package encryption

import (
	"Password-Keeper/pkg/etc"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"

	"golang.org/x/crypto/pbkdf2"
)

func EncryptData(data *etc.Data, plaintext string, masterKey string) error {
	salt := make([]byte, 16)
	var errSalt error
	_, errSalt = rand.Read(salt)
	if errSalt != nil {
		return fmt.Errorf("ошибка при генерации соли: %v", errSalt)
	}

	key := pbkdf2.Key([]byte(masterKey), salt, 4096, 32, sha256.New)

	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("ошибка при создании блочного шифра: %v", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("ошибка при работе с GCM: %v", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return fmt.Errorf("ошибка при генерации случайного значения шифрования: %v", err)
	}

	ciphertext := gcm.Seal(nil, nonce, []byte(plaintext), nil)

	data.EncryptedPassword = base64.StdEncoding.EncodeToString(ciphertext)
	data.Salt = base64.StdEncoding.EncodeToString(salt)
	data.Nonce = base64.StdEncoding.EncodeToString(nonce)

	return nil
}

func DecryptData(data *etc.Data, masterKey string) (string, error) {
	ciphertext, _ := base64.StdEncoding.DecodeString(data.EncryptedPassword)
	salt, _ := base64.StdEncoding.DecodeString(data.Salt)
	nonce, _ := base64.StdEncoding.DecodeString(data.Nonce)

	key := pbkdf2.Key([]byte(masterKey), salt, 4096, 32, sha256.New)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("ошибка при создании блочного шифра: %v", err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("ошибка при работе с GCM: %v", err)
	}
	if len(nonce) != gcm.NonceSize() {
		return "", fmt.Errorf("ошибка: некорректное значение длины случайного значения шифрования")
	}

	decryptedPassword, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("ошибка при расшифровке пароля: %v", err)
	}

	return string(decryptedPassword), nil
}
