package middlewares

import (
	"fmt"
	"net/http"

	"example.com/eventbrite/utils"
	"github.com/gin-gonic/gin"
)

func Aunthenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized, No token found"})
		return
	}

	userId, err := utils.VerifyToken(token)
	fmt.Println("RETURNED USER ID from authenticate function : ", userId)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
	}
	context.Set("userId", userId)
	context.Next()
}
