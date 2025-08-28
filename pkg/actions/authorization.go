package actions

import (
	"Password-Keeper/pkg/etc"
)

func (act *Action) AuthorizationUser() error {
	if err := act.auth.ReadConfig(&etc.Settings); err != nil {
		return err
	}

	user, err := act.auth.GetUserByUsername(etc.Settings.CurrentUsername)
	if err != nil {
		return err
	}

	if user == nil {
		if err := act.auth.SignUp(); err != nil {
			return err
		}
		return nil
	}

	for !act.auth.AllowedPass {
		if err := act.changeUser(); err != nil {
			return err
		}
	}

	return nil
}
