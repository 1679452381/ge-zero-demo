package logic

import (
	"context"
	user2 "go-zero-hello-2/mall/user/types/user"

	"go-zero-hello-2/mall/userapi/internal/svc"
	"go-zero-hello-2/mall/userapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserByIdLogic) GetUserById(req *types.IdRequest) (resp *types.UserResponse, err error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.UserRpc.GetUser(context.Background(), &user2.IdRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}
	resp = &types.UserResponse{
		Id:     user.Id,
		Name:   user.Name,
		Gender: user.Gender,
	}
	return
}
