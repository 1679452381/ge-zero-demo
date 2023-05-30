package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-hello-2/mall/user/internal/svc"
	"go-zero-hello-2/mall/user/model"
	"go-zero-hello-2/mall/user/types/user"
)

type UserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserLogic) GetUser(in *user.IdRequest) (*user.UserResponse, error) {

	res, err := l.svcCtx.UserRepo.FindOneById(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &user.UserResponse{
		Id:     res.Id,
		Name:   res.Name,
		Gender: res.Gender,
	}, nil
}

func (l *UserLogic) SaveUser(in *user.UserRequest) (*user.UserResponse, error) {
	data := &model.User{
		Name:   in.Name,
		Gender: in.Gender,
	}
	err := l.svcCtx.UserRepo.Save(l.ctx, data)
	if err != nil {
		return nil, err
	}
	return &user.UserResponse{
		Id:     data.Id,
		Name:   data.Name,
		Gender: data.Gender,
	}, nil
}

func (l *UserLogic) FindOneById(in *user.IdRequest) (*user.UserResponse, error) {
	// todo: add your logic here and delete this line

	return &user.UserResponse{}, nil
}
