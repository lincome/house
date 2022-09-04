package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AdminUsersModel = (*customAdminUsersModel)(nil)

type (
	// AdminUsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAdminUsersModel.
	AdminUsersModel interface {
		adminUsersModel
	}

	customAdminUsersModel struct {
		*defaultAdminUsersModel
	}
)

// NewAdminUsersModel returns a model for the database table.
func NewAdminUsersModel(conn sqlx.SqlConn, c cache.CacheConf) AdminUsersModel {
	return &customAdminUsersModel{
		defaultAdminUsersModel: newAdminUsersModel(conn, c),
	}
}
