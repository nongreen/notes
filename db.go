// db.go contains all functions for work with database

package main

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

// connectToDB abort app with panic if has errors
func connectToDB(config *mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return nil, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return nil, pingErr
	}

	return db, nil
}

// addUser adds user in db
func addUser(user *User, db *sql.DB) (int64, error) {
	userExisting, err := userIsExist(user, db)
	if err != nil {
		return 0, fmt.Errorf("addUser: %v", err)
	}
	if userExisting {
		return 0, errors.New("User is already exist")
	}

	result, err := db.Exec("INSERT INTO User (username, password, email) VALUES (?, ?, ?)", user.Username, user.Password, user.Email)
	if err != nil {
		return 0, fmt.Errorf("addUser: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addUser: %v", err)
	}

	return id, nil
}

// userIsExist returns True if is exist in db, and False if isn't
func userIsExist(user *User, db *sql.DB) (bool, error) {
	if user.ID == 0 {
		return false, nil
	}

	// Trying to get user from db
	var username, password, email string
	err := db.QueryRow("Select username, password, email from User where id = ?", user.ID).Scan(&username, &password, &email)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	// Check fields of user
	if user.Username != username {
		return false, nil
	}
	if user.Password != password {
		return false, nil
	}
	if user.Email != email {
		return false, nil
	}

	return true, nil
}
