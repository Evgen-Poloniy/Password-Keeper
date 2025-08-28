package actions

import (
	"Password-Keeper/pkg/etc"
	"encoding/json"
	"fmt"
	"os"
)

func (act *Action) saveAllData(filename string, data []etc.Data) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("ошибка открытия файла: %v", err)
	}
	defer file.Close()

	if len(data) == 0 {
		return nil
	}

	if _, err := file.WriteString("[\n  "); err != nil {
		return err
	}

	for i, item := range data {
		itemJSON, err := json.MarshalIndent(item, "  ", "  ")
		if err != nil {
			return fmt.Errorf("ошибка кодирования: %v", err)
		}

		if _, err := file.Write(itemJSON); err != nil {
			return err
		}

		if i < len(data)-1 {
			if _, err := file.WriteString(",\n  "); err != nil {
				return err
			}
		}
	}

	if _, err := file.WriteString("\n]"); err != nil {
		return err
	}

	return nil
}
