package actions

import (
	enc "Password-Keeper/pkg/encryption"
	"Password-Keeper/pkg/etc"
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/atotto/clipboard"
	"golang.org/x/crypto/bcrypt"
)

func (act *Action) copyTargetPassword() error {
	allData, err := act.readData()
	if err != nil {
		if err.Error() == "ошибка: файл не содержит данных" {
			fmt.Println("У вас нет сохраненных паролей для копирования")
			fmt.Println()
			etc.WaitInput()
			return nil
		}

		return err
	}

	data := make([]etc.Data, 0, len(allData))
	for _, dt := range allData {
		if dt.Username == etc.Settings.CurrentUsername {
			data = append(data, dt)
		}
	}

	if len(data) == 0 {
		fmt.Println("У вас нет сохраненных паролей для копирования")
		fmt.Println()
		etc.WaitInput()
		return nil
	}

	act.printAllData(data)

	scanner := bufio.NewScanner(os.Stdin)
	var passwordName string

	for {
		fmt.Println("Введите номер строки с паролем или название пароля, который хотите скопировать")
		scanner.Scan()
		passwordName = scanner.Text()

		etc.ClearConsole()

		if passwordName == "" {
			fmt.Println("Номер строки или имя пароля не должно быть пустым. Введите заного")
			fmt.Println()
			etc.WaitInput()
			continue
		}

		break
	}

	etc.ClearConsole()

	var decryptedPassword string
	positionUint64, err := strconv.ParseUint(passwordName, 10, 32)
	if err != nil {
		for _, dt := range data {
			if dt.PasswordName == passwordName {
				var masterKey string
				for {
					fmt.Printf("Введите мастер-пароль/пин-код от вашей учетной записи c именем пользователя \"%s\"\n", etc.Settings.CurrentUsername)
					fmt.Print("Пароль/пин-код: ")
					masterKey, err = etc.InputPassword()
					if err != nil {
						fmt.Printf("%v\n", err)
						fmt.Println()
						etc.WaitInput()
						continue
					}

					etc.ClearConsole()

					if err := bcrypt.CompareHashAndPassword([]byte(etc.Settings.CurrentUserPasswordHash), []byte(masterKey)); err != nil {
						fmt.Println("Пароли не совпадают")
						fmt.Println()
						etc.WaitInput()
						continue
					}

					break
				}

				decryptedPassword, err = enc.DecryptData(&dt, masterKey)
				if err != nil {
					return err
				}

				err = clipboard.WriteAll(decryptedPassword)
				if err != nil {
					return fmt.Errorf("ошибка при копировании пароля в буфер обмена: %v", err)
				}

				fmt.Printf("Пароль успешно скопирован и находится в буфере обмена.\nИспользуйте Ctrt + V для вставки в нужное место\n")
				fmt.Println()
				etc.WaitInput()

				break
			}
		}
	} else {
		position := int(positionUint64)
		for i, dt := range data {
			if position == i+1 {
				var masterKey string
				for {
					fmt.Printf("Введите мастер-пароль/пин-код от вашей учетной записи c именем пользователя \"%s\"\n", etc.Settings.CurrentUsername)
					fmt.Print("Пароль/пин-код: ")
					masterKey, err = etc.InputPassword()
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

					break
				}

				decryptedPassword, err = enc.DecryptData(&dt, masterKey)
				if err != nil {
					return err
				}

				err = clipboard.WriteAll(decryptedPassword)
				if err != nil {
					return fmt.Errorf("ошибка при копировании пароля в буфер обмена: %v", err)
				}

				fmt.Printf("Пароль успешно скопирован и находится в буфере обмена.\nИспользуйте Ctrt + V для вставки в нужное место\n")
				fmt.Println()
				etc.WaitInput()

				break
			}
		}
	}

	if decryptedPassword == "" {
		fmt.Println("Данной записи не существует")
		fmt.Println()
	}

	return nil
}
