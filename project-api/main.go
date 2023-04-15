package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/xszhangxiaocuo/CFC-MDP/project-api/api"
	"github.com/xszhangxiaocuo/CFC-MDP/project-api/config"
	"github.com/xszhangxiaocuo/CFC-MDP/project-api/router"
	srv "github.com/xszhangxiaocuo/CFC-MDP/project-common"
)

func main() {
	r := gin.Default()
	// 路由注册
	router.InitRouter(r)

	srv.Run(r, config.AppConf.ServerConfig.Name, config.AppConf.ServerConfig.Addr, nil)
}
