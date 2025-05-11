package main

import (
	"myapp/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.UserRoutes(r)
	r.Run()
}
