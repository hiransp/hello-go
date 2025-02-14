package main

import (
	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"ping": "pong",
	})
}

func main() {
	router := gin.Default()
	router.GET("/ping", ping)

	router.Run(":8000")
}
