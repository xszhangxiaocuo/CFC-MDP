package service

import "github.com/gin-gonic/gin"

/*
service模块路由处理函数
*/

// ServiceHandler 用于控制服务模块路由处理函数的结构体
type ServiceHandler struct {
}

func NewServiceHandler() *ServiceHandler {
	return &ServiceHandler{}
}

func (sh *ServiceHandler) index(ctx *gin.Context) {

}
