package main

import (
	"github.com/gin-gonic/gin"
	"github.com/skyclouds2001/SimpleDouyin/controller"
	"github.com/skyclouds2001/SimpleDouyin/service"
)

func main() {
	go service.RunMessageServer()

	r := gin.Default()

	r.Static("/static", "./public")

	router := r.Group("/douyin")

	// basic apis
	router.GET("/feed/", controller.Feed)
	router.GET("/user/", controller.UserInfo)
	router.POST("/user/register/", controller.Register)
	router.POST("/user/login/", controller.Login)
	router.POST("/publish/action/", controller.Publish)
	router.GET("/publish/list/", controller.PublishList)

	// extra apis - I
	router.POST("/favorite/action/", controller.FavoriteAction)
	router.GET("/favorite/list/", controller.FavoriteList)
	router.POST("/comment/action/", controller.CommentAction)
	router.GET("/comment/list/", controller.CommentList)

	// extra apis - II
	router.POST("/relation/action/", controller.RelationAction)
	router.GET("/relation/follow/list/", controller.FollowList)
	router.GET("/relation/follower/list/", controller.FollowerList)
	router.GET("/relation/friend/list/", controller.FriendList)
	router.GET("/message/chat/", controller.MessageChat)
	router.POST("/message/action/", controller.MessageAction)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
