package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	connect_to_db()
	router := gin.Default()
	router.LoadHTMLGlob("html/*.html")

	router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)

	})
	router.GET("/register", func(context *gin.Context) {
		context.HTML(http.StatusOK, "register_form.html", nil)

	})
	router.GET("/login", func(context *gin.Context) {
		context.HTML(http.StatusOK, "login_form.html", nil)

	})

	router.Run("localhost:8080")
}
