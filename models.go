package main

type User struct {
	ID       int64
	Username string
	Password string
	Email    string
}

func (user *User) is_valid() bool {
	if user.Username == "" {
		return false
	} else if user.Password == "" {
		return false
	} else if user.Email == "" {
		return false
	}
	return true
}

func (user *User) save() error {
	id, err := addUser(user, db)
	if err != nil {
		return err
	}
	user.ID = id

	return nil
}
