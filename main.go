package main

import (
	"SimpleDouyin/controller"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default()

	h.Static("/static", "./public")

	router := h.Group("/douyin")

	router.GET("/feed", controller.Feed)
	router.POST("/user/register", controller.Register)
	router.POST("/user/login", controller.Login)
	router.GET("/user", controller.UserInfo)
	router.POST("/publish/action")
	router.GET("/publish/list")

	h.Spin()
}
