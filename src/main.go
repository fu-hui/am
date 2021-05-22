package main

import (
	"github.com/fu-hui/am/src/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	userGroup := r.Group("/am/v1/user")

	// create user
	userGroup.POST("", controller.CreateUser)

	// user auth v1
	userGroup.POST("/auth/v1", controller.UserAuthV1)

	// service heartbeat test
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// listen and serve on 0.0.0.0:8080
	_ = r.Run()
}
