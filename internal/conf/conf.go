package conf

type ServerConf struct {
	Port          int  `mapstructure:"port"`
	EnableSwagger bool `mapstructure:"enable_swagger"`
}

type DBConf struct {
	DSN      string `mapstructure:"dsn"`
	LogLevel string `mapstructure:"log_level"`
}

type Conf struct {
	Debug  bool       `mapstructure:"debug"`
	Server ServerConf `mapstructure:"server"`
	DB     DBConf     `mapstructure:"db"`
}
