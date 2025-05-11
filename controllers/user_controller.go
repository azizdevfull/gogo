package controllers

import (
	"myapp/database"
	"myapp/models"
	"myapp/requests"
	"myapp/resources"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	c.JSON(http.StatusOK, resources.NewUserListResponse(users))
}

// find user by id
func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, resources.NewUserResponse(user))
}

func CreateUser(c *gin.Context) {
	var req requests.CreateUserRequest

	// Bind JSON body to struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Validate the request
	if err := requests.ValidateCreateUserRequest(req); err != nil {
		// Format errors
		var validationErrors []resources.ValidationError
		for _, e := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, resources.ValidationError{
				Field:   e.Field(),
				Message: e.Tag(),
			})
		}

		// Return custom error response
		c.JSON(http.StatusBadRequest, resources.ErrorResponse{
			Message:    "Validation failed",
			Validation: validationErrors,
		})
		return
	}

	// Proceed with creating the user
	user := models.User{Name: req.Name, Age: req.Age}
	database.DB.Create(&user)

	c.JSON(http.StatusCreated, user)
}

// Update user
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&user)
	c.JSON(http.StatusOK, resources.NewUserResponse(user))
}

// Delete user
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.User{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
