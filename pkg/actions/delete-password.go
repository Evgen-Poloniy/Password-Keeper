package actions

import (
	"Password-Keeper/pkg/etc"
	"bufio"
	"fmt"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func (act *Action) deletePassword() error {
	allData, err := act.readData()
	if err != nil {
		if err.Error() == "ошибка: файл не содержит данных" {
			fmt.Println("У вас нет сохраненных паролей для удаления")
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
		fmt.Println("У вас нет сохраненных паролей для удаления")
		fmt.Println()
		etc.WaitInput()
		return nil
	}

	scanner := bufio.NewScanner(os.Stdin)
	var passwordName string

	for {
		act.printAllData(data)
		fmt.Println()
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
		newData := make([]etc.Data, len(allData)-1)
		positionUint64, err := strconv.ParseUint(passwordName, 10, 32)
		if err != nil {
			var isExistRecord bool = false
			for _, dt := range data {
				if dt.PasswordName == passwordName {

					var i int = 0
					for _, t_dt := range allData {
						if t_dt.Username == etc.Settings.CurrentUsername {
							if t_dt.PasswordName == passwordName {
								isExistRecord = true
								continue
							}
						}

						newData[i] = dt
						i++
					}

				}
			}

			etc.ClearConsole()

			if !isExistRecord {
				fmt.Printf("Записи о пароле с именем \"%s\" не найдено\n", passwordName)
				fmt.Println()
				etc.WaitInput()
				continue
			}

			etc.ClearConsole()

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

			if err := act.makeBackup(etc.DatabaseFileName, allData); err != nil {
				fmt.Printf("ошибка при создания резервной копии базы данных \"%s\": %v\n", etc.DatabaseFileName, err)
				fmt.Println()
				etc.WaitInput()
				return nil
			}

			if err := act.saveAllData(etc.DatabaseFileName, newData); err != nil {
				fmt.Printf("ошибка при сохранении файла \"%s\": %v\n", etc.DatabaseFileName, err)
				fmt.Println()
				etc.WaitInput()
				return nil
			}

			fmt.Printf("Запись о пароле с названием \"%s\" успешно удалена\n", passwordName)
			fmt.Println()
			etc.WaitInput()

			break

		} else {
			position := int(positionUint64)
			if position > len(data) {
				fmt.Printf("Номер строки должен быть меньше %d\n", len(data)+1)
				fmt.Println()
				etc.WaitInput()
				continue
			} else if position < 1 {
				fmt.Println("Номер строки должен быть больше ли равен 1")
				fmt.Println()
				etc.WaitInput()
				continue
			}

			etc.ClearConsole()

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

			var targetData etc.Data
			for i, dt := range data {
				if i == position-1 {
					targetData = dt
				}
			}

			var i int = 0
			for _, dt := range allData {
				if dt.Username == etc.Settings.CurrentUsername {
					if dt == targetData {
						continue
					}
				}

				newData[i] = dt
				i++
			}

			if err := act.makeBackup(etc.DatabaseFileName, allData); err != nil {
				fmt.Printf("ошибка при создания резервной копии базы данных \"%s\": %v\n", etc.DatabaseFileName, err)
				fmt.Println()
				etc.WaitInput()
				return nil
			}

			if err := act.saveAllData(etc.DatabaseFileName, newData); err != nil {
				fmt.Printf("ошибка при сохранении файла \"%s\": %v\n", etc.DatabaseFileName, err)
				fmt.Println()
				etc.WaitInput()
				return nil
			}

			fmt.Printf("Запись о пароле с названием \"%s\" успешно удалена\n", passwordName)
			fmt.Println()
			etc.WaitInput()

		}

		break
	}

	etc.ClearConsole()

	return nil
}
