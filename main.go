package main

import (
	routes "event-booking/Routes"
	"event-booking/db"

	"github.com/gin-gonic/gin"
)

var Id = 1

func main() {
	server := gin.Default()
	db.DB = db.InitDB()
	db.CreateTables()
	routes.Routes(server)
	server.Run(":8080")
}
