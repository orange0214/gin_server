package global

import (
	"project/config"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	GVA_VP     *viper.Viper
	GVA_CONFIG config.Server
	GVA_LOG    *zap.Logger
)
