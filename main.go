package notes

import (
	"github.com/gin-gonic/gin"
)

func main() {
	db = connectToDB()

	router := gin.Default()
	router.LoadHTMLGlob("html/*.html")

	bindAllEndPoints(router)

	router.Run("localhost:8080")
}
