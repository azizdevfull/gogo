package main

import (
	"myapp/database"
	"myapp/models"
	"myapp/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.ConnectDB()
	database.DB.AutoMigrate(&models.User{})
	routes.UserRoutes(r)
	r.Run()
}
