package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/eventbrite/models"
	"example.com/eventbrite/utils"
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
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized, No token found"})
		return
	}

	err := utils.VerifyToken(token)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
	}

	var event models.Event
	err = context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request"})
	}
	id, err := event.SaveEvent()
	if err != nil {
		fmt.Println(err)
		return
	}
	event.Id = id
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

func UpdateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "cannot parse event id"})
		return
	}

	_, err = models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not fetch event"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Format"})
		return
	}

	updatedEvent.Id = eventId
	err = updatedEvent.UpdateEventById()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not update the field"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Updated Successfully"})
}

func DeleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "cannot parse event id"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Event doesnt Exist"})
		return
	}

	err = event.DeleteEventById()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to delete the event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event Deleted Successfully"})
}
