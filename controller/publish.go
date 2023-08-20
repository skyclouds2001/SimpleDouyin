package controller

import (
	"SimpleDouyin/common"
	"SimpleDouyin/repository"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type PublishActionResponse struct {
	common.Response
}

func PublishAction(_ context.Context, c *app.RequestContext) {
	token := c.PostForm("token")
	title := c.PostForm("title")
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(consts.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "",
		})
	}
	println(token, title, data)

	c.JSON(consts.StatusOK, PublishActionResponse{
		Response: common.Response{
			StatusCode: 0,
			StatusMsg:  "",
		},
	})
}

type PublishListResponse struct {
	common.Response
	VideoList []repository.Video `json:"video_list,omitempty"`
}

func PublishList(_ context.Context, c *app.RequestContext) {
	token := c.Query("token")
	userId := c.Query("user_id")
	println(userId, token)

	c.JSON(consts.StatusOK, PublishListResponse{
		Response: common.Response{
			StatusCode: 0,
			StatusMsg:  "",
		},
		VideoList: []repository.Video{},
	})
}
