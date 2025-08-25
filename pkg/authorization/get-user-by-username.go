package authorization

import (
	"Password-Keeper/pkg/etc"
	"encoding/json"
	"fmt"
	"os"
)

func (ath *Authorization) GetUserByUsername(targetUsername string) (*etc.User, error) {
	file, err := os.OpenFile(etc.UsersFileName, os.O_RDONLY, 0644)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, err
		}
		return nil, fmt.Errorf("ошибка открытия файла \"%s\": %v", etc.UsersFileName, err)
	}
	defer file.Close()

	var users []etc.User
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&users); err != nil {
		return nil, fmt.Errorf("ошибка чтения JSON-файла \"%s\": %v", etc.UsersFileName, err)
	}

	for _, user := range users {
		if user.Username == targetUsername {
			return &user, nil
		}
	}

	return nil, nil
}
