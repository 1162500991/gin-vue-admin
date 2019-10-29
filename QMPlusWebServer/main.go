package main

import (
	"QMPlusCommon/config"
	"QMPlusCommon/init/qmlog"
	"QMPlusCommon/init/qmsql"
	"fmt"
	"main/init/initRouter"
	"main/init/registTable"
	"net/http"
	"time"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	qmlog.InitLog()
	registTable.RegistTable(qmsql.InitMysql(config.Dbconfig.Admin))
	defer qmsql.DEFAULTDB.Close()
	Router := initRouter.InitRouter()
	//qmlog.QMLog.Info("服务器开启") // 日志测试代码

	s := &http.Server{
		Addr:           ":9999",
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	time.Sleep(10 * time.Microsecond)
	fmt.Printf(`欢迎使用 Gin-Vue-Admin
作者：奇淼 And Spike666
微信：shouzi_1994
默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
默认前端文件运行地址:http://127.0.0.1:8080
`, s.Addr)
	_ = s.ListenAndServe()
}
