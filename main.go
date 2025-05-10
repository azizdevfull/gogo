package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		name := c.Query("name")
		c.JSON(200, gin.H{
			"message": "pong " + name,
		})
	})

	r.Run()
}
