package user

import (
	"github.com/xszhangxiaocuo/CFC-MDP/project-common/discovery"
	"github.com/xszhangxiaocuo/CFC-MDP/project-grpc/user/login"
	"github.com/xszhangxiaocuo/CFC-MDP/project-user/config"
	"google.golang.org/grpc/resolver"
)

var LoginServiceClient login.LoginServiceClient

func InitRpcUserClient() {
	etcdRegister := discovery.NewResolver(config.AppConfig.EtcdConfig.Addrs)
	resolver.Register(etcdRegister)

}
