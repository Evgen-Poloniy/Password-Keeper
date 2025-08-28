package actions

import (
	"Password-Keeper/pkg/etc"
	"fmt"
	"os"

	"github.com/eiannone/keyboard"
)

var menuChangeUser = []string{
	"",
	"\n",
	"1   - Войти в текущую учетную запись",
	"2   - Войти в существующую учетную запись",
	"3   - Создать новую учетную запись",
	"\n",
	"Esc - выход",
}

func (act *Action) changeUser() error {
	var needRedrawMenu bool = true

	for {
		if needRedrawMenu {
			etc.ClearConsole()
			etc.PrintMenu(menuChangeUser)
		}

		needRedrawMenu = true

		err := keyboard.Open()
		if err != nil {
			etc.ClearConsole()
			return fmt.Errorf("ошибка открытия клавиатуры: %v", err)
		}

		char, key, _ := keyboard.GetKey()

		if key == keyboard.KeyEsc {
			keyboard.Close()
			etc.ClearConsole()

			if act.auth.IsFirstAuthorization {
				fmt.Println("Выход из программы...")
				os.Exit(0)
			}

			return nil
		}

		keyboard.Close()

		switch char {
		case '1':
			etc.ClearConsole()
			return act.auth.SignIn()

		case '2':
			etc.ClearConsole()
			act.auth.NeedInputUsername = true
			return act.auth.SignIn()

		case '3':
			etc.ClearConsole()
			return act.auth.SignUp()

		default:
			needRedrawMenu = false
			fmt.Printf("\a")
		}
	}
}
