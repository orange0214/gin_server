package config

type System struct {
	Addr         int    `mapstructure:"addr" json:"addr" yaml:"addr"`
	RouterPrefix string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
	DbType       string `mapstructure:"db-type" json:"db-type" yaml:"db-type"` // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
}
