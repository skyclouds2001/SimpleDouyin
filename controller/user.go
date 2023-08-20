package controller

import (
	"SimpleDouyin/common"
	"SimpleDouyin/repository"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type UserInfoResponse struct {
	common.Response
	User repository.User `json:"user,omitempty"`
}

func UserInfo(ctx context.Context, c *app.RequestContext) {
	userId := c.Query("user_id")
	token := c.Query("token")
	println(userId, token)
	c.JSON(consts.StatusOK, UserInfoResponse{
		Response: common.Response{
			StatusCode: 0,
			StatusMsg:  "",
		},
		User: repository.User{},
	})
}
