package router

import "github.com/gin-gonic/gin"

type Router interface { //路由接口
	Route(r *gin.Engine)
}

type RegisterRouter struct { //注册路由
}

var routers []Router

// InitRouter 批量注册路由
func InitRouter(r *gin.Engine) {
	for _, ro := range routers {
		ro.Route(r)
	}
}

func Register(ro ...Router) {
	routers = append(routers, ro...)
}
