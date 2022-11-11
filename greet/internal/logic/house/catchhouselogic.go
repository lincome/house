package house

import (
	"context"

	"go-zero-demo/greet/internal/svc"
	"go-zero-demo/greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CatchHouseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCatchHouseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CatchHouseLogic {
	return &CatchHouseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CatchHouseLogic) CatchHouse(req *types.CatchHouseReq) (resp *types.CatchHouseRes, err error) {
	return
}
