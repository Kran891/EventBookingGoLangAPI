package routes

import (
	"event-booking/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createUser(context *gin.Context) {
	var user models.User
	context.ShouldBindJSON(&user)
	user.Save()
	context.JSON(http.StatusCreated, user)
}
