package controllers

import (
	"fmt"
	"net/http"

	"example.com/eventbrite/models"
	"example.com/eventbrite/utils"
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
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the payload, send proper user info", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User Registered successfully", "user_id": userId})
}

func UserLogin(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the payload, send proper user info"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	token, err := utils.GenerateJwtToken(user.Email, user.Id)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
	}

	context.JSON(http.StatusOK, gin.H{"message": "Logged in Successfully", "token": token})
}
