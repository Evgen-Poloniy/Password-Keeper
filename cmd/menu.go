package main

import (
	"Password-Keeper/pkg/actions"
	"Password-Keeper/pkg/etc"
	"runtime"
)

var mainMenu = []string{
	"",
	"\n",
	"1   - сохранить новый пароль",
	"2   - скопировать пароль",
	"3   - удалить пароль",
	"\n",
	"4   - сменить пользователя",
	"\n",
	"Esc - выход",
}

type Menu struct {
	act *actions.Action
}

func NewMenu() *Menu {
	return &Menu{
		act: actions.NewAction(),
	}
}

func (m *Menu) redrawMenu() {
	if m.act.NeedRedrawMenu {
		etc.ClearConsole()
		etc.PrintMenu(mainMenu)
	}

	m.act.NeedRedrawMenu = true
}

func (m *Menu) checkColorsSupport() {
	if runtime.GOOS == "windows" {
		etc.ResetColors()
	}
}
