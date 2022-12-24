// models.go contains models for work with database and some method of it

package main

import "fmt"

type User struct {
	ID       int64
	Username string
	Password string
	Email    string
}

func (user *User) isValid() (bool, error) {
	usernameIsValid, err := user.usernameIsValid()
	if err != nil {
		return false, fmt.Errorf("UsernameIsValid function")
	}
	if !usernameIsValid {
		return false, nil
	}

	passwordIsValid, err := user.passwordIsValid()
	if err != nil {
		return false, fmt.Errorf("PasswordIsValid function")
	}
	if !passwordIsValid {
		return false, nil
	}

	emailIsValid, err := user.emailIsValid()
	if err != nil {
		return false, fmt.Errorf("EmailIsValid function")
	}
	if !emailIsValid {
		return false, nil
	}

	return true, nil
}

func (user *User) usernameIsValid() (bool, error) {
	if user.Username == "" {
		return false, nil
	}
	nonwordSymbolsInUsername, err := nonwordSymbolsInString(user.Username)
	if err != nil {
		return false, err
	}
	if user.Username == "" || nonwordSymbolsInUsername {
		return false, nil
	}
	return true, nil
}

func (user *User) passwordIsValid() (bool, error) {
	if user.Password == "" {
		return false, nil
	}
	passwordIsStrong, err := passwordHaveMinStrong(user.Password)
	if err != nil {
		return false, err
	}
	if !passwordIsStrong {
		return false, err
	}
	return true, nil
}

func (user *User) emailIsValid() (bool, error) {
	if user.Email == "" {
		return false, nil
	}
	emailIsEmail, err := stringIsEmail(user.Email)
	if err != nil {
		return false, err
	}
	if !emailIsEmail {
		return false, nil
	}
	return true, nil
}

func (user *User) save() error {
	id, err := addUser(user, db)
	if err != nil {
		return err
	}
	user.ID = id

	return nil
}
