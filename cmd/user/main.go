package main

import (
	"net"
	"strconv"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/schwarzsail/tiktok/config"
	"github.com/schwarzsail/tiktok/internal/user/adapter/db"
	"github.com/schwarzsail/tiktok/internal/user/adapter/rpc"
	service "github.com/schwarzsail/tiktok/internal/user/service/core"
	"github.com/schwarzsail/tiktok/pkg/kitex_gen/user/userservice"
	"github.com/schwarzsail/tiktok/pkg/logger"
)

var (
	etcdAddr           string
	serviceAddr        string
	serviceName        = "user"
	dbAdapter          *db.DBAdapter
	userServiceAdapter *service.CoreService
)

func init() {
	etcdAddr = config.GetEtcdAddr()
	serviceAddr = "0.0.0.0" + strconv.Itoa(config.GetAvailablePort())
	dbAdapter = db.NewDBAdapter()
	userServiceAdapter = service.NewCoreService(dbAdapter)
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{etcdAddr})
	if err != nil {
		logger.Fatalf("User: etcd registry failed, error: %v", err)
	}
	addr, err := net.ResolveTCPAddr("tcp", serviceAddr)
	if err != nil {
		logger.Fatalf("User: listen addr failed %v", err)
	}
	svr := userservice.NewServer(
		rpc.NewUserHandler(userServiceAdapter),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: serviceName,
		}),
		server.WithMuxTransport(),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
	)
	if err := svr.Run(); err != nil {
		logger.Fatalf("User: run server failed, error: %v", err)
	}
}
