package logic

import (
	"context"
	"go-zero-hello-2/mall/user/types/user"

	"go-zero-hello-2/mall/order/internal/svc"
	"go-zero-hello-2/mall/order/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GerOrderByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGerOrderByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GerOrderByIdLogic {
	return &GerOrderByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GerOrderByIdLogic) GerOrderById(req *types.OrderReq) (resp *types.OrderReply, err error) {
	// todo: add your logic here and delete this line
	//req.Id
	//建立grpc连接
	userResponse, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdRequest{Id: "1"})

	if err != nil {
		return nil, err
	}
	return &types.OrderReply{
		Id:       "1",
		Name:     "zxc",
		UserName: userResponse.Name,
	}, nil
}
