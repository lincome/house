package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ HouseModel = (*customHouseModel)(nil)

type (
	// HouseModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHouseModel.
	HouseModel interface {
		houseModel
	}

	customHouseModel struct {
		*defaultHouseModel
	}
)

// NewHouseModel returns a model for the database table.
func NewHouseModel(conn sqlx.SqlConn, c cache.CacheConf) HouseModel {
	return &customHouseModel{
		defaultHouseModel: newHouseModel(conn, c),
	}
}
