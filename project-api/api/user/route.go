package user

import (
	"github.com/gin-gonic/gin"
	"github.com/xszhangxiaocuo/CFC-MDP/project-api/router"
	"log"
)

// 初始化注册用户模块路由
func init() {
	log.Println("init user router")
	router.Register(&UserRouter{})
}

type UserRouter struct {
}

// Route 用户模块路由注册
func (*UserRouter) Route(r *gin.Engine) {
	//初始化grpc客户端连接
	InitRpcUserClient()
	uh := NewUserHandler()
	//为路由组注册中间件
	//group.Use(Auth())
	{
		r.POST("/login", uh.login)
	}
}
