package authorization

import (
	"Password-Keeper/pkg/etc"
	"bufio"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func (ath *Authorization) SignIn() error {
	scanner := bufio.NewScanner(os.Stdin)

	var user *etc.User
	var username string
	for {
		fmt.Print("Введите имя пользователя: ")
		scanner.Scan()
		username = scanner.Text()
		fmt.Println()

		if username == "" {
			fmt.Println("Имя пользователя не должно быть пустым")
			fmt.Println()
			etc.WaitInput()
			continue
		}

		var err error
		user, err = ath.GetUserByUsername(username)
		if err != nil {
			return err
		} else if user == nil {
			fmt.Println("Данной учетной записи не существует")
			fmt.Println()
			etc.WaitInput()
			continue
		}

		break
	}

	etc.ClearConsole()

	var password string
	for {
		var err error

		fmt.Printf("Введите пароль или пин-код от учетной записи \"%s\": ", username)
		password, err = etc.InputPassword()
		if err != nil {
			fmt.Printf("%v\n", err)
			fmt.Println()
			etc.WaitInput()
			continue
		}

		etc.ClearConsole()

		if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
			fmt.Println("Пароли не совпадают")
			fmt.Println()
			etc.WaitInput()
			continue
		}

		break
	}

	etc.Settings.CurrentUsername = user.Username
	etc.Settings.CurrentUserPasswordHash = user.PasswordHash

	ath.saveConfig(&etc.Settings)

	return nil
}
