package user

import (
	"github.com/xszhangxiaocuo/CFC-MDP/project-api/config"
	"github.com/xszhangxiaocuo/CFC-MDP/project-common/discovery"
	"github.com/xszhangxiaocuo/CFC-MDP/project-grpc/user/login"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"log"
)

var LoginServiceClient login.LoginServiceClient

func InitRpcUserClient() {
	//服务注册
	etcdRegister := discovery.NewResolver(config.AppConf.EtcdConfig.Addrs)
	resolver.Register(etcdRegister)
	conn, err := grpc.Dial("etcd:///user", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("can not connect: %v", err)
	}
	LoginServiceClient = login.NewLoginServiceClient(conn)
}
