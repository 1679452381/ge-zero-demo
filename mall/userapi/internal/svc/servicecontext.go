package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-hello-2/mall/user/userclient"
	"go-zero-hello-2/mall/userapi/internal/config"
	"go-zero-hello-2/mall/userapi/internal/middleware"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
	Example rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		Example: middleware.NewExampleMiddleware().Handle,
	}
}
