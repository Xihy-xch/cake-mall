package handler

import (
	"net/http"

	"cake-mall/api/internal/logic/member"
	"cake-mall/api/internal/svc"
	"cake-mall/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func UpdateWechatSessionKeyHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateWechatSessionKeyRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUpdateWechatSessionKeyLogic(r.Context(), ctx)
		resp, err := l.UpdateWechatSessionKey(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
