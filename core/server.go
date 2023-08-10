package core

import (
	"fmt"
	"project/global"
	"project/initialize"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))
	s := initServer(address, Router)
	fmt.Printf("address: %s\n", address)
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
