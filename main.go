package main

import (
	"gin-mongo-api/configs"
	"gin-mongo-api/routes/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"data": "Hello World",
		})
	})

	//mongoDB connect
	configs.ConnectDB()
	user.UserRoute(router)

	router.Run("localhost:6000")
}
