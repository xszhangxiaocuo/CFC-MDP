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
	//初始化grpc客户端连接
	InitRpcProjectClient()
	sh := NewServiceHandler()
	//为路由组注册中间件
	//group.Use(Auth())
	{
		r.POST("/new/save", sh.save)
	}
	group := r.Group("/list")
	{
		group.POST("/", sh.getList)
		group.POST("/send", sh.send)
		group.POST("/getTemplate", sh.getTemplate)
		group.POST("/modifyTemplate", sh.modifyTemplate)
	}
}
