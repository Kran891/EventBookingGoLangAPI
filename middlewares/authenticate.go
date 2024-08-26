package middlewares

import (
	"event-booking/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token, _ := context.Cookie("token")
	//token = strings.Split(token, " ")[1]
	fmt.Println(token)
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"msg": "UnAuthorized",
		})
	}
	id, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"msg": "UnAuthorized",
		})
	}
	context.Set("id", id)
	uid := context.GetInt64("id")
	fmt.Println(uid)
	context.Next()
}
