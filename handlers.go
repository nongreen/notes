// handlers.go contains handlers for all routers

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
	handleMailVerification(context)
}

// Responses login_form.html
func getLoginFormHandler(context *gin.Context) {
	context.HTML(http.StatusOK, "login_form.html", nil)
}
