package main

import (
	"github.com/gin-gonic/gin"
	srv "github.com/xszhangxiaocuo/CFC-MDP/project-common"
	"github.com/xszhangxiaocuo/CFC-MDP/project-user/config"
	"github.com/xszhangxiaocuo/CFC-MDP/project-user/router"
)

func main() {
	r := gin.Default()
	// 路由注册
	router.InitRouter(r)
	// grpc服务注册
	grpcServer := router.RegisterGrpc()
	// grpc服务注册到etcd
	router.RegisterEtcdServer()

	stop := func() {
		grpcServer.Stop()
	}
	srv.Run(r, config.AppConf.ServerConfig.Name, config.AppConf.ServerConfig.Addr, stop)
}
