package main

import (
	"flag"
	"fmt"
	"net/http"

	"go-zero-demo/common/errorx"
	"go-zero-demo/greet/global"
	"go-zero-demo/greet/internal/config"
	"go-zero-demo/greet/internal/handler"
	"go-zero-demo/greet/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/greet-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	// 自定义错误
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {
		case *errorx.CodeError:
			return http.StatusOK, e.Data()
		default:
			println(e.Error())
			return http.StatusInternalServerError, nil
		}
	})

	global.GVA_DB = global.Gorm(c) // gorm连接数据库

	fmt.Println("111")
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
