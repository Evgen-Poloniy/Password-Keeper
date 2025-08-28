package actions

import (
	"Password-Keeper/pkg/etc"
	"fmt"
)

func (act *Action) makeBackup(filename string, data []etc.Data) error {
	backupPath := fmt.Sprintf("%s.backup", filename)

	return act.saveAllData(backupPath, data)
}
