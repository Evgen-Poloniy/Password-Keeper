package authorization

import (
	"Password-Keeper/pkg/etc"
	"encoding/json"
	"fmt"
	"os"
)

func (ath *Authorization) saveConfig(settings *etc.UserSettings) error {
	file, err := os.OpenFile(etc.SettingsFileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("ошибка открытия файла: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(settings)
	if err != nil {
		return fmt.Errorf("ошибка записи в JSON-файл \"%s\": %v", etc.SettingsFileName, err)
	}

	return nil
}
