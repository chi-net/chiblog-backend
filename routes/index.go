package routes_index

import (
	"log"

	"github.com/gin-gonic/gin"
)

func route(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "200",
		"hello":  "World!",
	})
	log.Println("Requested at /")
}
