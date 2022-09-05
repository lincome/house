package global

import (
	"go-zero-demo/greet/internal/config"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
		DSN:                       c.Mysql1.DataSource, // DSN data source name
		DefaultStringSize:         191,                 // string 类型字段的默认长度
		SkipInitializeWithVersion: false,               // 根据版本自动配置
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,       // 禁用彩色打印
		},
	)

	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		Logger: newLogger,
	}); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		return db
	}
}
