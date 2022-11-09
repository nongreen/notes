package notes

import (
	"github.com/gin-gonic/gin"
)

func bindAllEndPoints(router *gin.Engine) {
	bindAllGET(router)
	router.POST("/register", postRegisterFormHandler)
}

func bindAllGET(router *gin.Engine) {
	router.GET("/", getIndexHandler)
	router.GET("/register", getRegisterFormHandler)
	router.GET("/login", getLoginFormHandler)
}
