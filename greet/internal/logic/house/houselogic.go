package house

import (
	"context"
	"strings"

	"go-zero-demo/common/errorx"
	"go-zero-demo/greet/global"
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
	groupString := ""
	groupNum := 10
	group := req.Group
	whereData := ""
	switch group {
	case "week":
		groupString = "left(date,10)"
		groupNum = 10
		whereData = carbon.Parse(carbon.Now().ToDateString()).SubDays(int(req.Limit)).ToDateTimeString()
	case "month":
		groupString = "left(date,7)"
		groupNum = 7
		whereData = carbon.Parse(carbon.Now().ToDateString()).SubMonthsNoOverflow(int(req.Limit)).ToDateTimeString()
	case "year":
		groupString = "left(date,4)"
		groupNum = 4
		whereData = carbon.Parse(carbon.Now().ToDateString()).SubYearsNoOverflow(int(req.Limit)).ToDateTimeString()
	default:
		groupString = "left(date,10)"
		groupNum = 10
		whereData = carbon.Parse(carbon.Now().ToDateString()).SubDays(int(req.Limit)).ToDateTimeString()
	}

	db := global.GVA_DB.Model(&model.HouseDb{})

	var total int64 = 0
	global.GVA_DB.Table("(?) as u", db.Model(&model.HouseDb{}).Select("id").Group(groupString).Group("area")).Count(&total)

	// pageSize := int(req.Limit) * 12
	// offset := int(total) - pageSize

	var houseData []model.House

	global.GVA_DB.Model(&model.HouseDb{}).Select("left(date,?) as date,area,sum(room_number) as room_number", groupNum).Where("date > ?", whereData).Group(groupString).Group("area").Order("date").Find(&houseData)

	// houseData, err := l.svcCtx.HouseModel.GetAllByArea(l.ctx, req.Area, req.Limit)
	// switch err {
	// case nil:
	// case model.ErrNotFound:
	// 	return nil, errorx.NewDefaultError("用户名不存在")
	// default:
	// 	return nil, err
	// }
	// println(len(houseData.List)) x轴
	var xData []string
	// println(len(houseData.List)) y轴

	areaRangeData := []string{"闽侯", "连江", "罗源", "闽清", "永泰", "长乐", "福清", "鼓楼区", "台江区", "晋安区", "马尾区", "仓山区"}

	var seriesData []types.SeriesData
	if len(houseData) > 0 {
		//x轴
		for _, homestayOrder := range houseData {
			xData = append(xData, homestayOrder.Date)
			// areaYData["id"] = append(areaYData["id"], homestayOrder.RoomNumber)
		}

		//组合区域map
		areaMap := make(map[string][]int64)
		for _, homestayOrder := range houseData {
			areaMap[homestayOrder.Area] = append(areaMap[homestayOrder.Area], homestayOrder.RoomNumber)
		}
		//组合series格式
		for _, areaData := range areaRangeData {
			var seriesDataData types.SeriesData
			seriesDataData.Data = areaMap[areaData]
			seriesDataData.Name = areaData
			seriesDataData.Type = "line"
			seriesData = append(seriesData, seriesDataData)
		}
	} else {
		return nil, errorx.NewDefaultError("没有数据")
	}
	var areaRange []types.AreaRange
	var limitRange []types.LimitRange
	var groupRange []types.GroupData

	limitRangeData := []int64{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 365}
	groupData := []string{"day", "month", "year"}
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
		XData: removeDuplication_map(xData),
		// YData:         yData,
		AreaRange:     areaRange,
		AreaRangeData: areaRangeData,
		LimitRange:    limitRange,
		GroupData:     groupRange,
		SeriesData:    seriesData,
	}, nil
}

func removeDuplication_map(arr []string) []string {
	set := make(map[string]struct{}, len(arr))
	j := 0
	for _, v := range arr {
		_, ok := set[v]
		if ok {
			continue
		}
		set[v] = struct{}{}
		arr[j] = v
		j++
	}

	return arr[:j]
}
