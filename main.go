package main

import (
	"net/http"

	"example.com/eventbrite/db"
	"example.com/eventbrite/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()
	server := gin.Default()

	server.GET("/events", getAllEvents)
	server.POST("/events", createEvent)

	server.Run(":8000")
}

func getAllEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request"})
	}

	event.Id = 1
	event.UserId = 10
	event.SaveEvent()
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}
