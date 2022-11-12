package house

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"go-zero-demo/greet/internal/svc"
	"go-zero-demo/greet/internal/types"
	"go-zero-demo/greet/model"

	"github.com/golang-module/carbon"
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
	// todo: add your logic here and delete this line
	// result, err := l.svcCtx.HouseModel.GetLastRecord(l.ctx)
	// if err != nil {
	// 	return nil, err
	// }

	date := carbon.Yesterday().ToDateString()
	if req.Date != "" {
		date = req.Date
	}

	println(date)

	result, err := l.GetHouseData(date)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	// fmt.Println("result:", result.Data.Pages)

	// l.svcCtx.HouseModel.GetRecordByDateAndArea(l.ctx)
	// if err != nil {
	// 	return nil, err
	// }

	j := 0
	for _, v := range result.DealNew {
		roomNumber, _ := strconv.Atoi(v.RoomNumber)
		//不存在更新
		d := &model.House{
			CreatedTime: carbon.Now().Carbon2Time(),
			Date:        carbon.Parse(date).Carbon2Time(),
			Area:        v.Area,
			RoomArea:    v.RoomArea,
			RoomNumber:  int64(roomNumber),
		}
		l.svcCtx.HouseModel.Insert(l.ctx, d)
		// if _, err = l.svcCtx.HouseModel.Insert(l.ctx, d); err != nil {
		// 	return l.CatchHouse(req)
		// }
		j++
	}

	for _, v := range result.DealOld {

		roomNumber, _ := strconv.Atoi(v.RoomNumber)
		//不存在更新
		d := &model.House{
			CreatedTime: carbon.Now().Carbon2Time(),
			Date:        carbon.Parse(date).Carbon2Time(),
			Area:        v.Area,
			RoomArea:    v.RoomArea,
			RoomNumber:  int64(roomNumber),
		}
		l.svcCtx.HouseModel.Insert(l.ctx, d)
		// if _, err = l.svcCtx.HouseModel.Insert(l.ctx, d); err != nil {
		// 	return l.CatchHouse(req)
		// }
		j++
	}
	return
}

func (l *CatchHouseLogic) GetHouseData(date string) (result House, err error) {
	spaceClient := http.Client{
		Timeout: time.Second * 10, // Maximum of 2 secs
		Transport: &http.Transport{
			// 设置代理
			// Proxy:           http.ProxyURL(uri),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	getUrl := l.svcCtx.Config.HouseApi + date[0:10]
	req, err := http.NewRequest(http.MethodGet, getUrl, nil)
	if err != nil {
		return
	}

	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	// req.Header.Set("Accept-Encoding", "gzip, deflate")
	// req.Header.Set("Content-Type", "application/json; charset=utf-8")

	res, getErr := spaceClient.Do(req)

	if getErr != nil {
		err = getErr
		return
	}
	defer res.Body.Close()
	fmt.Println("status", res.Status)

	body, readErr := ioutil.ReadAll(res.Body)
	// fmt.Println(string(body))
	if readErr != nil {
		err = readErr
		return
	}
	jsonErr := json.Unmarshal(body, &result)
	if jsonErr != nil {
		err = jsonErr
		return
	}
	// fmt.Println(result)
	// fmt.Println("result:", result.Data.Pages)
	return
}

type House struct {
	Date    string    `json:"date"`
	DealNew []DealNew `json:"deal_new"`
	Sale    []Sale    `json:"sale"`
	Left    []Left    `json:"left"`
	DealOld []DealOld `json:"deal_old"`
}
type DealNew struct {
	Area       string `json:"area"`
	RoomNumber string `json:"room_number"`
	RoomArea   string `json:"room_area"`
}
type List struct {
	HomeOuterName string      `json:"home_outer_name"`
	SpRoom        json.Number `json:"sp_room"`
	PjRoom        json.Number `json:"pj_room"`
}
type Sale struct {
	Area string      `json:"area"`
	Sum  json.Number `json:"sum"`
	List []List      `json:"list"`
}
type Left struct {
	Area       string `json:"area"`
	RoomNumber string `json:"room_number"`
	RoomArea   string `json:"room_area"`
	HomeNumber string `json:"home_number"`
}
type DealOld struct {
	Area       string `json:"area"`
	RoomNumber string `json:"room_number"`
	RoomArea   string `json:"room_area"`
}
