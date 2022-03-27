package main

import "github.com/gin-gonic/gin"

func main() {
	app := gin.Default()
	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "200",
			"hello":  "World!",
		})
	})
	app.Run()
}
