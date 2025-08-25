package authorization

import (
	"Password-Keeper/pkg/etc"
	"encoding/json"
	"fmt"
	"os"
)

func (ath *Authorization) ReadConfig(settings *etc.UserSettings) error {
	file, err := os.OpenFile(etc.SettingsFileName, os.O_RDONLY, 0644)
	if err != nil {
		if os.IsNotExist(err) {
			return err
		}
		return fmt.Errorf("ошибка открытия файла \"%s\": %v", etc.SettingsFileName, err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(settings)
	if err != nil {
		return fmt.Errorf("ошибка чтения JSON-файла \"%s\": %v", etc.SettingsFileName, err)
	}

	return nil
}
