package main

import (
	"os"

	"github.com/go-sql-driver/mysql"
)

var DBCONFIG mysql.Config = mysql.Config{
	User:                 os.Getenv("DBUSER"),
	Passwd:               os.Getenv("DBPASS"),
	Net:                  "tcp",
	Addr:                 "127.0.0.1:3306",
	DBName:               "notes",
	AllowNativePasswords: true,
}
