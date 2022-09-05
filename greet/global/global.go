package global

import (
	"go-zero-demo/greet/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	GVA_DB     *gorm.DB
	GVA_DBList map[string]*gorm.DB
)

func Gorm(c config.Config) *gorm.DB {
	return GormMysql(c)
}

// GormMysql 初始化Mysql数据库
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func GormMysql(c config.Config) *gorm.DB {
	// var c config.Config
	// m := global.GVA_CONFIG.Mysql
	// if m.Dbname == "" {
	// 	return nil
	// }
	mysqlConfig := mysql.Config{
		DSN:                       "c.Mysql1.DataSource", // DSN data source name
		DefaultStringSize:         191,                   // string 类型字段的默认长度
		SkipInitializeWithVersion: false,                 // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig)); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		return db
	}
}
