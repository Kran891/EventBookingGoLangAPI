package routes

import "github.com/gin-gonic/gin"

func Routes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.POST("/events", postEvent)
	server.PUT("/events/", updateEvent)
	server.GET("/events/:id", getEvent)
	server.DELETE("events/:id", deleteEvent)
	server.POST("/users", createUser)
	server.PUT("/users", updateUser)
	server.GET("/users/:id", getUser)
	server.DELETE("/users/:id", deleteUser)
}
