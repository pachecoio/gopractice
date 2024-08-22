package exercices

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sayHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}

func StartWebServer() {
	router := gin.Default()
	router.GET("/", sayHello)

	router.Run()
}
