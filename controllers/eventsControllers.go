package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/eventbrite/models"
	"github.com/gin-gonic/gin"
)

func GetAllEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	context.JSON(http.StatusOK, events)
}

func CreateEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request"})
	}

	event.SaveEvent()
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func GetEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "cannot parse event id"})
		return
	}

	event, err := models.GetEventById(eventId)
	fmt.Println("errors1 is", err)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Event not found in the database"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": event})
}
