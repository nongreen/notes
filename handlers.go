package main

import (
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
		return
	}

	if !(user.is_valid()) {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := user.save()
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
