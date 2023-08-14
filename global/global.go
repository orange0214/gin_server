package global

import (
	"project/config"
	"sync"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GVA_VP     *viper.Viper
	GVA_CONFIG config.Server
	GVA_LOG    *zap.Logger
	GVA_DB     *gorm.DB
	GVA_DBList map[string]*gorm.DB
	lock       sync.RWMutex
)

// 通过数据库名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return GVA_DBList[dbname]
}

func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := GVA_DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}
