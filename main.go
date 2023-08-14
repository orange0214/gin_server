package main

import (
	"project/core"
	"project/global"
	"project/initialize"

	"go.uber.org/zap"
)

func main() {
	// 配置管理
	global.GVA_VP = core.Viper()
	// 日志管理
	global.GVA_LOG = core.Zap()
	zap.ReplaceGlobals(global.GVA_LOG)
	core.RunWindowsServer()
	// 连接至数据库
	global.GVA_DB = initialize.Gorm()
	initialize.DBList()

}
