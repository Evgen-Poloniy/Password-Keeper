package authorization

import (
	"Password-Keeper/pkg/etc"
	"bufio"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func (ath *Authorization) SignUp() error {
	var user etc.User

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Придумайте и введите имя пользователя: ")
		scanner.Scan()
		user.Username = scanner.Text()
		fmt.Println()

		if user.Username == "" {
			fmt.Println("Имя пользователя не должно быть пустым")
			fmt.Println()
			etc.WaitInput()
			continue
		}

		break
	}

	etc.ClearConsole()

	for {
		fmt.Print("Придумайте простой пароль (от 4 до 16 символов) или пин-код: ")
		password, err := etc.InputPassword()
		if err != nil {
			fmt.Printf("%v\n", err)
			fmt.Println()
			etc.WaitInput()
			continue
		}

		etc.ClearConsole()

		var length int = len(password)
		if length < 4 || length > 16 {
			fmt.Printf("ошибка: пароль должен содержать от 4 до 16 символов.\nТекущий пароль содержит %d символов\n", length)
			fmt.Println()
			etc.WaitInput()
			continue
		}

		bytePasswordHash, errGenerateHash := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if errGenerateHash != nil {
			fmt.Printf("ошибка генерации хэша пароля: %v", errGenerateHash)
			fmt.Println()
			etc.WaitInput()
			continue
		}

		user.PasswordHash = string(bytePasswordHash)
		etc.Settings.CurrentUsername = user.Username
		etc.Settings.CurrentUserPasswordHash = user.PasswordHash

		ath.saveConfig(&etc.Settings)

		break
	}

	ath.saveAtFile(etc.UsersFileName, &user)

	ath.AllowedPass = true
	ath.IsFirstAuthorization = false

	return nil
}
