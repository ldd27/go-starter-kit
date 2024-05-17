package setup

import (
	"github.com/ldd27/go-starter-kit/internal/conf"
	"github.com/ldd27/go-starter-kit/pkg/zapx"
	"go.uber.org/zap"
)

func SetupZAP(conf conf.Conf) error {
	opt := func(option *zapx.Option) {
		option.Debug = conf.Debug
	}
	logger, err := zapx.New(opt)
	if err != nil {
		return err
	}

	zap.RedirectStdLog(logger)
	zap.ReplaceGlobals(logger)
	return nil
}
