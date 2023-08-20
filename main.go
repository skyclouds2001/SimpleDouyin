package main

import (
	"SimpleDouyin/controller"
	"SimpleDouyin/repository"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	repository.Init()

	h := server.Default()

	h.Static("/static", "./public")

	router := h.Group("/douyin")

	router.GET("/feed", controller.Feed)
	router.POST("/user/register", controller.Register)
	router.POST("/user/login", controller.Login)
	router.GET("/user", controller.UserInfo)
	router.POST("/publish/action", controller.PublishAction)
	router.GET("/publish/list", controller.PublishList)

	h.Spin()
}
