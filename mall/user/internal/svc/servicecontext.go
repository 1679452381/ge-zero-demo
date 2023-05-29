package svc

import (
	"go-zero-hello-2/mall/user/database"
	"go-zero-hello-2/mall/user/internal/config"
	"go-zero-hello-2/mall/user/internal/dao"
	"go-zero-hello-2/mall/user/repo"
)

type ServiceContext struct {
	Config   config.Config
	UserRepo repo.UserRpo
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := database.ConnectDb(c.Mysql.DataSource)
	return &ServiceContext{
		Config:   c,
		UserRepo: dao.NewUserDao(conn),
	}
}
