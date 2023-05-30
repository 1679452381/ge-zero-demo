package logic

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"go-zero-hello-2/common/errorx"
	"go-zero-hello-2/mall/user/model"
	"go-zero-hello-2/mall/user/types/user"
	"net/http"
	"strconv"
	"strings"
	"time"

	"go-zero-hello-2/mall/userapi/internal/svc"
	"go-zero-hello-2/mall/userapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginQequest) (resp *types.LoginReply, err error) {
	if len(strings.TrimSpace(req.Id)) == 0 {
		return nil, errors.New("参数错误")
	}

	userInfo, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdRequest{Id: req.Id})
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errorx.NewDefaultCodeError("用户不存在")
	default:
		return nil, errorx.NewDefaultCodeError("用户不存在")
	}
	id, err := strconv.Atoi(userInfo.Id)
	if err != nil {
		return nil, err
	}
	// ---start---
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, int64(id))
	if err != nil {
		return nil, err
	}
	// ---end---

	return &types.LoginReply{
		StatusCode: http.StatusOK,
		Message:    "success",
		Token: types.TokenInfo{
			AccessToken:  jwtToken,
			AccessExpire: now + accessExpire,
			RefreshAfter: now + accessExpire/2,
		},
	}, nil
}

// 生成token
func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
