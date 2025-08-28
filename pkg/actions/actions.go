package actions

import (
	auth "Password-Keeper/pkg/authorization"
	"Password-Keeper/pkg/etc"
	"fmt"
	"os"

	"github.com/eiannone/keyboard"
)

type Action struct {
	NeedRedrawMenu bool
	auth           *auth.Authorization
}

func NewAction() *Action {
	return &Action{
		NeedRedrawMenu: true,
		auth:           auth.NewAuthorization(),
	}
}

func (act *Action) ExecuteAction() error {
	err := keyboard.Open()
	if err != nil {
		return fmt.Errorf("ошибка открытия клавиатуры: %v", err)
	}

	char, key, _ := keyboard.GetKey()

	if key == keyboard.KeyEsc {
		keyboard.Close()
		etc.ClearConsole()
		fmt.Println("Выход из программы...")
		os.Exit(0)
	}

	keyboard.Close()

	switch char {
	case '1':
		etc.ClearConsole()
		err = act.saveNewPassword()
	case '2':
		etc.ClearConsole()
		err = act.copyTargetPassword()
	case '3':
		etc.ClearConsole()
		// err = act.deletePassword()
		// fmt.Println("Функция удаления еще не реализована до конца. Чтобы удалить пароль, найдите файл \"database.json\"")
		// fmt.Printf("по пути \"%s\" из корневой папки программы и удалите запись вручную.\n", etc.DatabaseFileName)
		// fmt.Println("Помните, что в JSON-формате данные отделены {}, поэтому нужно удалить нужный блок со скобками,")
		// fmt.Println("а разделетельную запятую так же убрать!")
		// fmt.Println()
		// etc.WaitInput()

	case '4':
		etc.ClearConsole()
		err = act.changeUser()

	default:
		act.NeedRedrawMenu = false
		fmt.Printf("\a")
	}

	return err
}
