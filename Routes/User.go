package routes

import (
	"event-booking/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func createUser(context *gin.Context) {
	var user models.User
	context.ShouldBindJSON(&user)
	user.Save()
	context.JSON(http.StatusCreated, user)
}
func getUser(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "Please give Id as a number",
		})
		return
	}
	user := models.GetUserById(id)
	context.JSON(http.StatusOK, user)

}
func updateUser(context *gin.Context) {
	var user models.User
	context.ShouldBindJSON(&user)
	user.Update()
	context.JSON(http.StatusAccepted, user)
}
