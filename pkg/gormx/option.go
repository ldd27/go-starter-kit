package gormx

type Option struct {
	DSN           string `json:"dsn"`
	LogLevel      string `json:"log_level"`
	Colorful      bool   `json:"colorful"`
	SlowThreshold int    `json:"slow_threshold"` // 单位ms
}

var defaultOption = Option{
	LogLevel:      "info",
	Colorful:      true,
	SlowThreshold: 200,
}
