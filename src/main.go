package main

import (
	"github.com/fu-hui/am/src/controller"
	"github.com/fu-hui/am/src/dao"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	// 1. init db
	if err := dao.InitDb(); err != nil {
		log.Printf("init db fail, err is:%v\n", err)
		os.Exit(1)
	}
	log.Println("init db service success")

	// 2. init web
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

	go func() {
		// listen and serve on 0.0.0.0:8080
		if err := r.Run(); err != nil {
			log.Printf("init web fail, err is:%v\n", err)
			os.Exit(1)
		}
	}()

	log.Println("am service start end")
	select {}
}
