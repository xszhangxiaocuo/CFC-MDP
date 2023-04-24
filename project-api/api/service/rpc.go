package service

import (
	"github.com/xszhangxiaocuo/CFC-MDP/project-api/config"
	"github.com/xszhangxiaocuo/CFC-MDP/project-common/discovery"
	projectRpc "github.com/xszhangxiaocuo/CFC-MDP/project-grpc/service/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"log"
)

var ProjectServiceClient projectRpc.ProjectServiceClient

func InitRpcProjectClient() {
	etcdRegister := discovery.NewResolver(config.AppConf.EtcdConfig.Addrs)
	resolver.Register(etcdRegister)
	conn, err := grpc.Dial("etcd:///service", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("can not connect:%v", err)
	}
	ProjectServiceClient = projectRpc.NewProjectServiceClient(conn)
}
