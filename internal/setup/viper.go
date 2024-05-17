package setup

import (
	"strings"

	"github.com/iamolegga/enviper"
	"github.com/ldd27/go-starter-kit/internal/constant"
	"github.com/ldd27/go-starter-kit/internal/setup/g"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func SetupViper(confPath string) error {
	e := enviper.New(viper.New())
	e.AutomaticEnv()
	e.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	e.SetEnvPrefix(constant.EnvPrefix)

	if confPath != "" {
		e.SetConfigFile(confPath)
		if err := e.MergeInConfig(); err != nil {
			zap.L().Error("merge config err", zap.Error(err))
			return err
		}
	}

	if err := e.Unmarshal(&g.C); err != nil {
		zap.L().Error("unmarshal err", zap.Error(err))
		return err
	}

	return nil
}
