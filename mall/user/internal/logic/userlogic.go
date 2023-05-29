package logic

import (
	"context"
	"go-zero-hello-2/mall/user/internal/svc"
	"go-zero-hello-2/mall/user/model"
	"go-zero-hello-2/mall/user/types/user"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
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
	// todo: add your logic here and delete this line

	return &user.UserResponse{
		Id:     in.GetId(),
		Name:   "zxc",
		Gender: "1",
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
		Id:     strconv.Itoa(int(data.Id)),
		Name:   data.Name,
		Gender: data.Gender,
	}, nil
}
