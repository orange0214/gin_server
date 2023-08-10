package main

import (
	"project/core"
	"project/global"

	"go.uber.org/zap"
)

func main() {
	// 配置管理
	global.GVA_VP = core.Viper()
	// 日志管理
	global.GVA_LOG = core.Zap()
	zap.ReplaceGlobals(global.GVA_LOG)
	core.RunWindowsServer()
}
