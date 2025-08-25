package actions

import (
	"encoding/json"
	"fmt"
	"os"
)

func (act *Action) saveAtFile(filename string, data interface{}) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("ошибка открытия файла: %v", err)
	}
	defer file.Close()

	fileInfo, _ := os.Stat(filename)

	if fileInfo.Size() > 0 {
		file.Seek(-2, os.SEEK_END)
		file.WriteString(",\n  ")
	} else {
		file.WriteString("[\n  ")
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("  ", "  ")

	err = encoder.Encode(data)
	if err != nil {
		fmt.Printf("ошибка записи в JSON-файл \"%s\": %v\n", filename, err)
	}

	file.WriteString("]")

	return nil
}
