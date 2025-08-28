package actions

import (
	"Password-Keeper/pkg/etc"
	"encoding/json"
	"fmt"
	"os"
)

func (act *Action) readData() ([]etc.Data, error) {
	fileInfo, err := os.Stat(etc.DatabaseFileName)
	if os.IsNotExist(err) {
		fmt.Println("Вы еще не сохранили ни одного пароля")
		return nil, fmt.Errorf("ошибка: %v", err)
	}

	if fileInfo.Size() == 0 {
		return nil, fmt.Errorf("ошибка: файл не содержит данных")
	}

	file, err := os.OpenFile(etc.DatabaseFileName, os.O_RDONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия файла \"%s\": %v", etc.DatabaseFileName, err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var data []etc.Data
	err = decoder.Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения JSON-файла \"%s\": %v", etc.DatabaseFileName, err)
	}

	return data, nil
}
