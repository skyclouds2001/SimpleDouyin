package controller

import (
	"SimpleDouyin/common"
	"SimpleDouyin/repository"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type PublishListResponse struct {
	common.Response
	VideoList []repository.Video `json:"video_list"`
}

func PublishList(ctx context.Context, c *app.RequestContext) {
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
