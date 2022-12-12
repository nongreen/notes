/* change var to const */
package main

import (
	"os"

	"github.com/go-sql-driver/mysql"
)

var SECRETKEY string = os.Getenv("SECRETKEY")

var GIN_MODE = os.Getenv("GIN_MODE")

var DBCONFIG mysql.Config = mysql.Config{
	User:                 os.Getenv("DBUSER"),
	Passwd:               os.Getenv("DBPASS"),
	Net:                  "tcp",
	Addr:                 "127.0.0.1:3306",
	DBName:               "notes",
	AllowNativePasswords: true,
}

// Server settings
const DOMAIN string = "localhost"
const PORT string = "8080"

// Validators settings
const USERNAMELEN = 64
const PASSLEN = 256
const EMAIL = 256

var USERKEY = "user"
