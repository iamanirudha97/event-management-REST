package main

import (
	"example.com/eventbrite/db"
	"example.com/eventbrite/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()
	server := gin.Default()
	routes.EventRoutes(server)
	server.Run(":8000")
}
