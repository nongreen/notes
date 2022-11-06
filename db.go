package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func connectToDB() *sql.DB {
	db, errors := sql.Open("mysql", DBCONFIG.FormatDSN())
	if errors != nil {
		panic(errors)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	return db
}

func addUser(user *User) (int64, error) {
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
