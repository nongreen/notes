package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	db, _ = connectToDB(&DBCONFIG)

	router := gin.Default()
	router.LoadHTMLGlob("html/*.html")

	bindAllEndPoints(router)

	router.Run("localhost:8080")
}
