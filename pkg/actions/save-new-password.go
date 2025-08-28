package actions

import (
	enc "Password-Keeper/pkg/encryption"
	"Password-Keeper/pkg/etc"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/eiannone/keyboard"
	"golang.org/x/crypto/bcrypt"
)

func (act *Action) saveNewPassword() error {
	var (
		passwordName string
		description  string
		password     string
	)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Введите название пароля (краткое название ресурса, для которого это пароль используется)")
		fmt.Println("Внимание! Не пишите на этом этапе пароль, т.к. ввод будет осуществлен позже")
		scanner.Scan()
		passwordName = scanner.Text()
		if passwordName == "" {
			etc.ClearConsole()
			fmt.Println("Название пароля не должно быть пустым")
			fmt.Println()
			etc.WaitInput()
			continue
		}

		_, err := strconv.ParseUint(passwordName, 10, 32)
		if err == nil {
			etc.ClearConsole()
			fmt.Println("Название пароля не может содержать только числа!")
			fmt.Println("Введите название пароля, в котором будет хотя бы одна буква")
			fmt.Println()
			etc.WaitInput()
			continue
		}

		etc.ClearConsole()

		break
	}

	fmt.Println("Напишите описание пароля (необязательно - можно оставить пустым)")
	fmt.Println("Внимание! Не пишите на этом этапе пароль, т.к. вы уже ввели его")

	scanner.Scan()
	description = scanner.Text()

	etc.ClearConsole()

	for {
		fmt.Printf("Оставьте пустым, чтобы сгенерировать пароль длинной %d символов\nВведите новый пароль (от 8-и до 16 символов): ", etc.Settings.PasswordGenerationLenth)
		var err error
		password, err = etc.InputPassword()
		if err != nil {
			fmt.Printf("%v\n", err)
			fmt.Println()
			etc.WaitInput()
			continue
		}

		etc.ClearConsole()

		if password == "" {
			fmt.Println("Вы действительно уверены, что хотите сгенерировать случайный пароль?")
			fmt.Println()
			fmt.Println("Нажмите Enter, чтобы подтвердить, либо Esc, чтобы отменить")

			err := keyboard.Open()
			if err != nil {
				etc.ClearConsole()
				fmt.Printf("ошибка открытия клавиатуры: %v", err)
				fmt.Println()
				etc.WaitInput()
				continue
			}

			_, key, _ := keyboard.GetKey()

			keyboard.Close()

			switch key {
			case keyboard.KeyEnter:
				password, err = enc.GeneratePassword(etc.Settings.PasswordGenerationLenth)
				if err != nil {
					fmt.Printf("%v\n", err)
					fmt.Println()
					etc.WaitInput()
					continue
				}

			case keyboard.KeyEsc:
				continue
			}

			etc.ClearConsole()
		}

		var length int = len(password)
		if length < 8 || length > 16 {
			fmt.Printf("Пароль должен содержать от 8 до 16 символов.\nТекущий пароль содержит %d символов\n", length)
			fmt.Println()
			etc.WaitInput()
			continue
		}

		break
	}

	data := etc.Data{
		Username:     etc.Settings.CurrentUsername,
		PasswordName: passwordName,
		Description:  description,
	}

	for {
		fmt.Printf("Введите мастер-пароль/пин-код от вашей учетной записи c именем пользователя \"%s\"\n", etc.Settings.CurrentUsername)
		fmt.Print("Пароль/пин-код: ")

		masterKey, err := etc.InputPassword()
		if err != nil {
			fmt.Printf("%v\n", err)
			fmt.Println()
			etc.WaitInput()
			continue
		}

		etc.ClearConsole()

		if errHash := bcrypt.CompareHashAndPassword([]byte(etc.Settings.CurrentUserPasswordHash), []byte(masterKey)); errHash != nil {
			fmt.Println("Пароли не совпадают")
			fmt.Println()
			etc.WaitInput()
			continue
		}

		if err = enc.EncryptData(&data, password, masterKey); err != nil {
			fmt.Printf("%v\n", err)
			continue
		}

		break
	}

	etc.ClearConsole()

	data.DateOfCreation = time.Now().Format("2006-01-02 15:04:05")

	if err := act.addDataAtFile(etc.DatabaseFileName, data); err != nil {
		fmt.Printf("ошибка при сохранении файла \"%s\": %v", etc.DatabaseFileName, err)
		fmt.Println()
		etc.WaitInput()
		return nil
	}

	fmt.Println("Пароль успешно сохранен")
	fmt.Println()
	etc.WaitInput()

	return nil
}
