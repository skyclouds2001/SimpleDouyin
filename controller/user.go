package controller

import (
	"SimpleDouyin/common"
	"SimpleDouyin/repository"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type RegisterResponse struct {
	common.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}

func Register(_ context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")
	println(username, password)

	c.JSON(consts.StatusOK, RegisterResponse{
		Response: common.Response{
			StatusCode: 0,
			StatusMsg:  "",
		},
		UserId: 0,
		Token:  "",
	})
}

type LoginResponse struct {
	common.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}

func Login(_ context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")
	println(username, password)

	c.JSON(consts.StatusOK, LoginResponse{
		Response: common.Response{
			StatusCode: 0,
			StatusMsg:  "",
		},
		UserId: 0,
		Token:  "",
	})
}

type UserInfoResponse struct {
	common.Response
	User *repository.User `json:"user,omitempty"`
}

func UserInfo(_ context.Context, c *app.RequestContext) {
	userId := c.Query("user_id")
	token := c.Query("token")
	println(userId, token)

	c.JSON(consts.StatusOK, UserInfoResponse{
		Response: common.Response{
			StatusCode: 0,
			StatusMsg:  "",
		},
		User: &repository.User{},
	})
}
