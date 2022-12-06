package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getIndexHandler(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", nil)
}

func getRegisterFormHandler(context *gin.Context) {
	context.HTML(http.StatusOK, "register_form.html", nil)
}

func postRegisterFormHandler(context *gin.Context) {
	user := &User{}
	if err := context.Bind(user); err != nil {
		return
	}

	if userIsValid, err := user.isValid(); !userIsValid || err != nil {
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

func getLoginFormHandler(context *gin.Context) {
	context.HTML(http.StatusOK, "login_form.html", nil)
}
