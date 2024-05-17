package setup

import (
	"github.com/ldd27/go-starter-kit/internal/conf"
	"github.com/ldd27/go-starter-kit/internal/model"
	"github.com/ldd27/go-starter-kit/internal/setup/g"
	"github.com/ldd27/go-starter-kit/pkg/gormx"
	"go.uber.org/zap"
)

func SetupDB(conf conf.Conf) error {
	db, err := gormx.New(func(option *gormx.Option) {
		option.DSN = conf.DB.DSN
		option.LogLevel = conf.DB.LogLevel
	})
	if err != nil {
		zap.L().Error("new db err", zap.Error(err))
		return err
	}

	g.SetDB(db)
	return nil
}

func AutoMigrate(conf conf.Conf) error {
	tables := []interface{}{
		&model.Example{},
	}

	if err := g.DB().AutoMigrate(tables...); err != nil {
		zap.L().Error("auto migrate err", zap.Error(err))
		return err
	}

	return nil
}
