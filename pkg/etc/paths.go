package etc

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	SettingsFileName string
	UsersFileName    string
	DatabaseFileName string
)

func GetPaths() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("ошибка при чтении .env файла: %v", err)
	}

	SettingsFileName = os.Getenv("PASSWORD_KEEPER_SETTINGS")
	UsersFileName = os.Getenv("PASSWORD_KEEPER_USERS")
	DatabaseFileName = os.Getenv("PASSWORD_KEEPER_DATABASE")

	if SettingsFileName == "" || UsersFileName == "" || DatabaseFileName == "" {
		return fmt.Errorf("ошибка при чтении .env файла: не удалось прочесть переменные окружения")
	}

	return nil
}
