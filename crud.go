package main

import (
	"github.com/gin-gonic/gin"
	"web_service/initializers"
)

func init() {
	initializers.LoadVaribles()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
