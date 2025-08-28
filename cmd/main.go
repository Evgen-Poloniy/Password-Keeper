package main

import (
	"Password-Keeper/pkg/etc"
	"fmt"
)

func main() {
	// ath := auth.NewAuthorization()
	// act := actions.NewAction(ath)
	// menu := NewMenu(act)

	etc.ClearConsole()

	menu := NewMenu()

	menu.checkColorsSupport()

	if err := menu.act.AuthorizationUser(); err != nil {
		panic(err)
	}

	// ctx, cancel := context.WithCancel(context.Background())

	for {
		menu.redrawMenu()

		if err := menu.act.ExecuteAction(); err != nil {
			fmt.Printf("%s%v%s\n", etc.RedColor, err, etc.WhiteColor)
			fmt.Println()
			etc.WaitInput()
		}
	}
}
