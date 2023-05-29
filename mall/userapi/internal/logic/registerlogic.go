package logic

import (
	"context"
	"go-zero-hello-2/mall/user/types/user"
	"time"

	"go-zero-hello-2/mall/userapi/internal/svc"
	"go-zero-hello-2/mall/userapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.UserRegisterQequest) (resp *types.Response, err error) {
	//超时上下文 超时自动断开
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	data := &user.UserRequest{Name: req.Name, Gender: req.Gender}
	_, err = l.svcCtx.UserRpc.SaveUser(ctx, data)
	if err != nil {
		resp = &types.Response{
			Message:    "服务器错误",
			StatusCode: 500,
		}
		return
	}

	resp = &types.Response{
		Message:    "添加成功",
		StatusCode: 200,
		Data:       data,
	}
	return
}
