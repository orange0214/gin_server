package config

type Server struct {
	System System          `mapstructure:"system" json:"system" yaml:"system"`
	Zap    Zap             `mapstructure:"zap" json:"zap" yaml:"zap"`
	Mysql  Mysql           `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	DBList []SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`
}
