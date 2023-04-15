package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/xszhangxiaocuo/CFC-MDP/project-api/api/pkg/model/user"
	"github.com/xszhangxiaocuo/CFC-MDP/project-grpc/user/login"
	"net/http"
	"time"

	common "github.com/xszhangxiaocuo/CFC-MDP/project-common"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (uh *UserHandler) login(ctx *gin.Context) {
	// 1. 接受参数 参数模型
	resp := &common.Result{}
	var req user.LoginReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.Fail(http.StatusBadRequest, "参数错误"))
		return
	}
	// 2.调用user grpc 完成登录
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	rpcReq := &login.LoginRequest{
		Account:  req.Account,
		Password: req.Password,
	}
	loginRpcResp, err := LoginServiceClient.Login(c, rpcReq)
	if err != nil {
		ctx.JSON(http.StatusOK, err)
		return
	}
	_ = loginRpcResp
	ctx.JSON(http.StatusOK, resp.Success(""))
}
