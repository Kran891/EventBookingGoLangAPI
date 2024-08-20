package routes

import (
	"event-booking/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, models.GetAllEvents())
}
func postEvent(context *gin.Context) {
	var event models.Event
	context.ShouldBindJSON(&event)
	event.Save()
	context.JSON(http.StatusCreated, event)
}
func updateEvent(context *gin.Context) {
	var event models.Event
	context.ShouldBindJSON(&event)
	event.Update()
	context.JSON(http.StatusAccepted, event)
}
func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	fmt.Println(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "Please check the Id Provided"})
	}
	event := models.GetById(id)
	context.JSON(http.StatusOK, event)
}
func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "Please check the Id Provided"})
	}
	models.DeleteById(id)
	context.JSON(http.StatusOK, id)
}
