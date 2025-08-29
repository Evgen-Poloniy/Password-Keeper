package main

import (
	"Password-Keeper/pkg/etc"
	"fmt"
	"os"
)

func main() {
	etc.ClearConsole()

	if err := etc.GetPaths(); err != nil {
		fmt.Printf("%v\n", err)
		fmt.Println()
		etc.WaitInput()
		os.Exit(1)
	}

	menu := NewMenu()

	if err := menu.act.AuthorizationUser(); err != nil {
		fmt.Printf("ошибка при авторизации пользователя: %v\n", err)
		fmt.Println()
		etc.WaitInput()
		os.Exit(1)
	}

	for {
		menu.redrawMenu()

		if err := menu.act.ExecuteAction(); err != nil {
			fmt.Printf("%v\n", err)
			fmt.Println()
			etc.WaitInput()
		}
	}
}
