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
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	r.POST("/", func(c *gin.Context) {
		name := c.PostForm("name")
		if name == "" {
			c.JSON(400, gin.H{
				"error": "name is required",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Received",
			"name":    name,
		})
	})

	r.PUT("/edit/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"message": "Edit",
			"id":      id,
		})
	})

	r.DELETE("/delete/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"message": "Delete",
			"id":      id,
		})
	})

	r.Run()
}
