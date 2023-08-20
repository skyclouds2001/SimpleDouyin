package controller

import (
	"SimpleDouyin/common"
	"SimpleDouyin/repository"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"time"
)

type FeedResponse struct {
	common.Response
	VideoList []repository.Video `json:"video_list,omitempty"`
	NextTime  int64              `json:"next_time,omitempty"`
}

func Feed(ctx context.Context, c *app.RequestContext) {
	c.JSON(consts.StatusOK, FeedResponse{
		Response: common.Response{
			StatusCode: 0,
			StatusMsg:  "",
		},
		VideoList: []repository.Video{},
		NextTime:  time.Now().Unix(),
	})
}
