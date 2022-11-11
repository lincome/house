package house

import (
	"go-zero-demo/greet/internal/logic/house"
	"go-zero-demo/greet/internal/svc"
	"go-zero-demo/greet/internal/types"
	"go-zero-demo/greet/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func HouseHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HouseReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := house.NewHouseLogic(r.Context(), svcCtx)
		resp, err := l.House(&req)
		response.Response(w, resp, err)
	}
}
