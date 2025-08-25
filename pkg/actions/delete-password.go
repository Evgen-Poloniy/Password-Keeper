package actions

import (
	"Password-Keeper/pkg/etc"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func (act *Action) deletePassword() error {
	allData, err := act.readData()
	if err != nil {
		return err
	}

	data := make([]etc.Data, 0, len(allData))
	for _, dt := range allData {
		if dt.Username == etc.Settings.CurrentUsername {
			data = append(data, dt)
		}
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

		positionUint64, err := strconv.ParseUint(passwordName, 10, 32)
		if err != nil {
			var isExistRecord bool = false
			for _, dt := range data {
				if dt.PasswordName == passwordName {
					// newData := make([]etc.Data, len(allData))
					var newData []etc.Data
					for _, t_dt := range allData {
						if t_dt.Username == etc.Settings.CurrentUsername {
							if t_dt.PasswordName == passwordName {
								continue
							}
						}

						newData = append(newData, t_dt)
					}

					if err := act.clearFile(etc.DatabaseFileName); err != nil {
						fmt.Printf("ошибка при очистке файла: %v\n", err)
						fmt.Println()
						etc.WaitInput()
						return nil
					}

					if err := act.saveAtFile(etc.DatabaseFileName, newData); err != nil {
						fmt.Printf("ошибка при сохранении файла \"%s\": %v\n", etc.DatabaseFileName, err)
						fmt.Println()
						etc.WaitInput()
						return nil
					}

					fmt.Printf("Запись о пароле с названием \"%s\" успешно удалена\n", passwordName)
					fmt.Println()
					etc.WaitInput()

					return nil
				}
			}

			etc.ClearConsole()

			if !isExistRecord {
				fmt.Printf("Записи о пароле с именем \"%s\" не найдено\n", passwordName)
				fmt.Println()
				etc.WaitInput()
				continue
			}

		} else {
			position := int(positionUint64)
			if position > len(data) {
				fmt.Printf("Номер строки должен быть меньше %d\n", len(data))
				fmt.Println()
				etc.WaitInput()
				continue
			} else if position < 1 {
				fmt.Println("Номер строки должен быть больше 1")
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

			// newData := make([]etc.Data, len(allData))
			var newData []etc.Data
			for _, dt := range allData {
				if dt.Username == etc.Settings.CurrentUsername {
					if dt == targetData {
						continue
					}
				}

				newData = append(newData, dt)
			}

			if err := act.clearFile(etc.DatabaseFileName); err != nil {
				fmt.Printf("ошибка при очистке файла: %v\n", err)
				fmt.Println()
				etc.WaitInput()
				return nil
			}

			if err := act.saveAtFile(etc.DatabaseFileName, newData); err != nil {
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
