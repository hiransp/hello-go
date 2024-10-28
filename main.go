package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func main() {
	fmt.Println("hello")

	router := gin.Default()
	router.GET("/", sayHello)
	router.Run()
}
