package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Responses index.html
func getIndexHandler(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", nil)
}

// Responses register_form.html
func getRegisterFormHandler(context *gin.Context) {
	context.HTML(http.StatusOK, "register_form.html", nil)
}

// Trying save uesr from form
func postRegisterFormHandler(context *gin.Context) {
	user := &User{}
	if err := context.Bind(user); err != nil {
		if GIN_MODE == "debug" {
			log.Println("Error bind")
			log.Fatal(err)
		}
		return
	}
	userIsValid, err := user.isValid()
	if err != nil {
		if GIN_MODE == "debug" {
			log.Println("Error userIsValid")
			log.Fatal(err)
		}
	}
	if !userIsValid {
		if GIN_MODE == "debug" {
			log.Println("User is valid")
		}
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = user.save()
	if err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	/* todo: redirect to page with message */
	context.Redirect(http.StatusFound, "/")
}

// Responses login_form.html
func getLoginFormHandler(context *gin.Context) {
	context.HTML(http.StatusOK, "login_form.html", nil)
}
