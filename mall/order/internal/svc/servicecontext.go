package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-hello-2/mall/order/internal/config"
	"go-zero-hello-2/mall/user/userclient"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
