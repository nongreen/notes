package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	var err error
	db, err = connectToDB(&DBCONFIG)

	if err != nil {
		log.Fatal("Cannot connect to test db")
		os.Exit(1)
	}

	router := gin.Default()
	router.LoadHTMLGlob("html/*.html")

	bindAllEndPoints(router)

	router.Run("localhost:8080")
}
