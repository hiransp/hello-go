package main

import (
	"github.com/gin-gonic/gin"
)

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func main() {

	router := gin.Default()
	router.GET("/", sayHello)
	router.Run()
}
