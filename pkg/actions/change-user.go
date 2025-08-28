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
	"",
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

			if etc.Settings.CurrentUsername == "admin" {
				menuChangeUser[2] = ""
			} else {
				if !act.auth.AllowedPass {
					menuChangeUser[2] = "1   - Войти в текущую учетную запись"
				} else {
					menuChangeUser[2] = ""
				}
			}

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
			// if !act.auth.AllowedPass {
			// 	etc.ClearConsole()
			// 	return act.auth.SignIn()
			// }

			if etc.Settings.CurrentUsername == "admin" {
				needRedrawMenu = false
				fmt.Printf("\a")
			} else {
				if !act.auth.AllowedPass {
					etc.ClearConsole()
					return act.auth.SignIn()
				} else {
					needRedrawMenu = false
					fmt.Printf("\a")
				}
			}

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
