package user

import (
	"github.com/gin-gonic/gin"
	"github.com/xszhangxiaocuo/CFC-MDP/project-api/router"
	"log"
)

// 初始化注册用户模块路由
func init() {
	log.Println("init user router")
	router.Register()
}

type UserRouter struct {
}

// Route 用户模块路由注册
func (*UserRouter) Route(r *gin.Engine) {
	uh := NewUserHandler()
	//注册路由组
	group := r.Group("/CFC-MDP")
	//为路由组注册中间件
	//group.Use(Auth())
	{
		group.POST("/login", uh.login)

	}
}
