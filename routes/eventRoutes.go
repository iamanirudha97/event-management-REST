package routes

import (
	"example.com/eventbrite/controllers"
	"example.com/eventbrite/middlewares"
	"github.com/gin-gonic/gin"
)

func EventRoutes(server *gin.Engine) {
	server.GET("/events", controllers.GetAllEvents)
	server.GET("/events/:id", controllers.GetEvent)

	authRoutes := server.Group("/")
	authRoutes.Use(middlewares.Aunthenticate)
	authRoutes.POST("/events", controllers.CreateEvent)
	authRoutes.PUT("/events/:id", controllers.UpdateEvent)
	authRoutes.DELETE("/events/:id", controllers.DeleteEvent)
	authRoutes.POST("/events/:id/register", controllers.RegisterForEvent)
	authRoutes.DELETE("/events/:id/register", controllers.CancelUserRegistration)

	server.POST("/signup", controllers.UserRegister)
	server.POST("/login", controllers.UserLogin)
}
