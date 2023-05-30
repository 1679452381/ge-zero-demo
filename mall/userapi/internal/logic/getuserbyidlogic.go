package logic

import (
	"context"
	"fmt"
	"go-zero-hello-2/mall/user/types/user"

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

func (l *GetUserByIdLogic) GetUserById(req *types.IdRequest) (resp *types.Response, err error) {

	user, err := l.svcCtx.UserRpc.GetUser(context.Background(), &user.IdRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}
	fmt.Printf("id:%s\n", req.Id)

	fmt.Printf("user:%s,err:%v\n", user, err)
	resp = &types.Response{
		Message:    "success",
		StatusCode: 200,
		Data:       user,
	}
	return
}
