package zapx

import (
	"context"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Option struct {
	Debug     bool
	zapOption []zap.Option
}

func (r *Option) With(opts ...zap.Option) {
	r.zapOption = append(r.zapOption, opts...)
}

var (
	defaultOption = Option{}
)

func New(opts ...func(option *Option)) (logger *zap.Logger, err error) {
	opt := defaultOption
	for _, o := range opts {
		o(&opt)
	}

	var config zap.Config
	if opt.Debug {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		config = zap.NewProductionConfig()
	}

	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	config.Level = zap.NewAtomicLevel()
	if opt.Debug {
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}
	config.DisableStacktrace = true

	return config.Build(opt.zapOption...)
}

func Logger(ctx context.Context, opts ...zap.Option) *zap.Logger {
	reqID, ok := ctx.Value(echo.HeaderXRequestID).(string)
	if ok {
		return zap.L().WithOptions(opts...).With(zap.Any("request_id", reqID))
	} else {
		return zap.L().WithOptions(opts...)
	}
}
