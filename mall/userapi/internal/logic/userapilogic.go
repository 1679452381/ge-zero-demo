package logic

import (
	"context"
	"go-zero-hello-2/mall/user/types/user"
	"go-zero-hello-2/mall/userapi/internal/svc"
	"go-zero-hello-2/mall/userapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserapiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserapiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserapiLogic {
	return &UserapiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserapiLogic) Userapi(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	data := &user.UserRequest{
		Name:   "zxczxc",
		Gender: "男",
	}
	_, err = l.svcCtx.UserRpc.SaveUser(l.ctx, data)
	if err != nil {
		resp = &types.Response{Message: "出错了", StatusCode: 500}
		return
	}
	resp = &types.Response{Message: "添加成功", StatusCode: 200}
	return
}
