package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", HelloIndex)
	router.GET("/ping", HelloWeb)
	router.POST("/ping", HelloWeb)
	router.Run()
}

//HelloWeb init point
func HelloWeb(c *gin.Context) {

	if c.Request.Method == "POST" {

		if value, ok := c.GetPostForm("name"); ok {
			log.Println("FOUNDDDDDDD IT", value)
		}

		c.HTML(http.StatusOK, "internal.html", nil)
		return
	}

	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func HelloIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
