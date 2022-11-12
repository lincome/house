package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ HouseModel = (*customHouseModel)(nil)

type (
	// HouseModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHouseModel.
	HouseModel interface {
		houseModel
		GetRecordByDate(ctx context.Context, date string) error
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

func (c *customHouseModel) GetRecordByDate(ctx context.Context, date string) error {
	m := &House{}
	query := fmt.Sprintf("select * from %s where `date` = ? limit 1", c.table)
	err := c.CachedConn.QueryRowNoCacheCtx(ctx, m, query, date)
	return err
}
