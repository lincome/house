package handler

import (
	"go-zero-demo/greet/internal/logic"
	"go-zero-demo/greet/internal/svc"
	"go-zero-demo/greet/internal/types"
	"go-zero-demo/greet/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func houseHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HouseReq
		if err := httpx.Parse(r, &req); err != nil {
			// httpx.Error(w, err)
			response.Response(w, nil, err)
			return
		}

		l := logic.NewHouseLogic(r.Context(), svcCtx)
		resp, err := l.House(&req)
		response.Response(w, resp, err)
	}
}
