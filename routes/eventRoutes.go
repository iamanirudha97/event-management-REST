package routes

import (
	"example.com/eventbrite/controllers"
	"github.com/gin-gonic/gin"
)

func EventRoutes(server *gin.Engine) {
	server.GET("/events", controllers.GetAllEvents)
	server.GET("/events/:id", controllers.GetEvent)
	server.POST("/events", controllers.CreateEvent)
	server.PUT("/events/:id", controllers.UpdateEvent)
	server.DELETE("/events/:id", controllers.DeleteEvent)
}
