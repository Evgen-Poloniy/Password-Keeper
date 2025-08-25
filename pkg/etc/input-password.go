package etc

import (
	"fmt"
	"syscall"

	"golang.org/x/term"
)

func InputPassword() (string, error) {
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", fmt.Errorf("ошибка чтения введенного пароля: %v", err)
	}

	password := string(bytePassword)
	fmt.Println()

	return password, nil
}
