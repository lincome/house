package logic

import (
	"context"
	"strings"

	"go-zero-demo/common/errorx"
	"go-zero-demo/greet/internal/svc"
	"go-zero-demo/greet/internal/types"
	"go-zero-demo/greet/model"

	"github.com/golang-module/carbon"
	"github.com/zeromicro/go-zero/core/logx"
)

type HouseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHouseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HouseLogic {
	return &HouseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HouseLogic) House(req *types.HouseReq) (resp *types.HouseReply, err error) {
	// todo: add your logic here and delete this line
	if len(strings.TrimSpace(req.Area)) == 0 {
		return nil, errorx.NewDefaultError("参数错误")
	}

	houseData, err := l.svcCtx.HouseModel.GetAllByArea(l.ctx, req.Area, req.Limit)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errorx.NewDefaultError("用户名不存在")
	default:
		return nil, err
	}
	// println(len(houseData.List)) x轴
	var xData []string
	// println(len(houseData.List)) y轴
	var yData []int64
	if len(houseData) > 0 {
		for _, homestayOrder := range houseData {
			xData = append(xData, carbon.Parse(homestayOrder.Date).ToDateString())
			yData = append(yData, homestayOrder.RoomNumber)
		}
	} else {
		return nil, errorx.NewDefaultError("没有数据")
	}
	var areaRange []types.AreaRange
	var limitRange []types.LimitRange
	var groupRange []types.GroupData

	areaRangeData := []string{"闽侯", "连江", "罗源", "闽清", "永泰", "长乐", "福清", "鼓楼区", "台江区", "晋安区", "马尾区", "仓山区"}
	limitRangeData := []int64{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 365}
	groupData := []string{"day", "week", "month", "year"}
	for _, area := range areaRangeData {
		areaRange = append(areaRange, types.AreaRange{
			Label: area,
			Value: area,
		})
	}
	for _, limit := range limitRangeData {
		limitRange = append(limitRange, types.LimitRange{
			Label: limit,
			Value: limit,
		})
	}
	for _, group := range groupData {
		groupRange = append(groupRange, types.GroupData{
			Label: group,
			Value: group,
		})
	}

	return &types.HouseReply{
		XData:      xData,
		YData:      yData,
		AreaRange:  areaRange,
		LimitRange: limitRange,
		GroupData:  groupRange,
	}, nil
}
