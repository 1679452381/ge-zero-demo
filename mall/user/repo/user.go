package repo

import (
	"context"
	"go-zero-hello-2/mall/user/model"
)

// 接口 在internal 的dao中实现逻辑
type UserRpo interface {
	Save(ctx context.Context, user *model.User) error
	FindOneById(ctx context.Context, id string) (*model.User, error)
}
