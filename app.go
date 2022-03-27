package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "200",
			"hello":  "World!",
		})
		log.Println("Requested at /")
	})
	app.Run()
}
