package svc

import (
	"go-zero-demo/greet/internal/config"
	"go-zero-demo/greet/internal/middleware"
	"go-zero-demo/greet/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config     config.Config
	UserModel  model.AdminUsersModel
	HouseModel model.HouseModel
	Example    rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	conn1 := sqlx.NewMysql(c.Mysql1.DataSource)
	return &ServiceContext{
		Config:     c,
		UserModel:  model.NewAdminUsersModel(conn, c.CacheRedis),
		HouseModel: model.NewHouseModel(conn1, c.CacheRedis),
		Example:    middleware.NewExampleMiddleware().Handle,
	}
}
