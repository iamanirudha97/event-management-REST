package controllers

import (
	"fmt"
	"net/http"

	"example.com/eventbrite/models"
	"github.com/gin-gonic/gin"
)

func UserRegister(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the payload, send proper user info"})
		return
	}

	userId, err := user.SaveUser()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the payload, send proper user info", "error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User Registered successfully", "user_id": userId})
}
