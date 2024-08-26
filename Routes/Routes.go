package routes

import (
	"event-booking/middlewares"

	"github.com/gin-gonic/gin"
)

func Routes(server *gin.Engine) {
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	server.GET("/events", getEvents)
	authenticated.POST("/events", postEvent)
	authenticated.PUT("/events/", updateEvent)
	authenticated.DELETE("events/:id", deleteEvent)
	server.GET("/events/:id", getEvent)

	server.POST("/users", createUser)
	server.PUT("/users", updateUser)
	server.GET("/users/:id", getUser)
	server.DELETE("/users/:id", deleteUser)
	server.POST("/login", loginUser)
}
