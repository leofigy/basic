package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", HelloIndex)
	router.GET("/ping", HelloWeb)
	router.Run()
}

//HelloWeb init point
func HelloWeb(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func HelloIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Index hello",
	})
}
