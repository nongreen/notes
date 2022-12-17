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

// SMTP server
var SMTP_HOST = os.Getenv("SMTP_HOST")
var SMTP_PORT = os.Getenv("SMTP_PORT")
var SMTP_ACCOUNT = os.Getenv("SMTP_ACCOUNT")
var SMTP_PASSWORD = os.Getenv("SMTP_PASSWORD")

// Server settings
const DOMAIN = "localhost"
const PORT = "8080"

// Validators settings
const USERNAMELEN = 64
const PASSLEN = 256
const EMAILLEN = 256

var USERKEY = "user"
