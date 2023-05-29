package handler

import (
	"go-zero-hello-2/mall/userapi/internal/logic"
	"go-zero-hello-2/mall/userapi/internal/svc"
	"go-zero-hello-2/mall/userapi/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserapiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserapiLogic(r.Context(), svcCtx)
		resp, err := l.Userapi(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
