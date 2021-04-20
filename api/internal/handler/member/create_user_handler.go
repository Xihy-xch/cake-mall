package handler

import (
	"net/http"

	"cake-mall/api/internal/logic/member"
	"cake-mall/api/internal/svc"
	"cake-mall/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func CreateUserHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateUserRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCreateUserLogic(r.Context(), ctx)
		resp, err := l.CreateUser(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
