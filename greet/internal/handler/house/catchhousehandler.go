package house

import (
	"go-zero-demo/greet/internal/logic/house"
	"go-zero-demo/greet/internal/svc"
	"go-zero-demo/greet/internal/types"
	"go-zero-demo/greet/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CatchHouseHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CatchHouseReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := house.NewCatchHouseLogic(r.Context(), svcCtx)
		resp, err := l.CatchHouse(&req)
		response.Response(w, resp, err)
	}
}
