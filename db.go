package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// connectToDB abort app with panic if has errors
func connectToDB() *sql.DB {
	db, errors := sql.Open("mysql", DBCONFIG.FormatDSN())
	if errors != nil {
		panic(errors)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		panic(errors)
	}

	return db
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

	result, err := db.Exec("Select id, username from User where id = ?", user.ID)
	if err != nil {
		return false, err
	}
	fmt.Println(result)
	return true, nil
}
