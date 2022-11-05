package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func connect_to_db() {
	db, errors := sql.Open("mysql", DBCONFIG.FormatDSN())
	if errors != nil {
		panic(errors)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}
