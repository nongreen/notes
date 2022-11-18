package notes

import (
	"github.com/gin-gonic/gin"
)

func bindAllEndPoints(router *gin.Engine) {
	router.GET("/", getIndexHandler)

	router.GET("/register", getRegisterFormHandler)
	router.POST("/register", postRegisterFormHandler)

	router.GET("/login", getLoginFormHandler)
}
