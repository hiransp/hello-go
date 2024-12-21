package main

import (
	"github.com/gin-gonic/gin"
)

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"ping": "pong",
	})
}

func main() {
	router := gin.Default()
	router.GET("/", sayHello)
	router.GET("/ping", ping)

	router.Run()
}
