package service

import (
	"github.com/gin-gonic/gin"
	"github.com/xszhangxiaocuo/CFC-MDP/project-api/router"
	"log"
)

// 初始化注册服务模块路由
func init() {
	log.Println("init service router")
	router.Register(&ServiceRouter{})
}

type ServiceRouter struct {
}

// Route 服务模块路由注册
func (*ServiceRouter) Route(r *gin.Engine) {
	sh := NewServiceHandler()
	//注册路由组
	group := r.Group("/CFC-MDP")
	//为路由组注册中间件
	//group.Use(Auth())
	{
		group.POST("/index", sh.index)

	}
}
