package gormx

import (
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	PostgresPrefixDSN = "postgres://"
	MysqlPrefixDSN    = "mysql://"
)

func New(opts ...func(option *Option)) (*gorm.DB, error) {
	opt := defaultOption
	for _, v := range opts {
		v(&opt)
	}

	newLogger := NewLogger(opt)
	switch opt.LogLevel {
	case "error":
		newLogger = newLogger.LogMode(logger.Error)
	case "warn":
		newLogger = newLogger.LogMode(logger.Warn)
	default:
		newLogger = newLogger.LogMode(logger.Info)
	}
	gormConf := &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	var dialector gorm.Dialector

	if strings.HasPrefix(opt.DSN, MysqlPrefixDSN) {
		opt.DSN = strings.ReplaceAll(opt.DSN, MysqlPrefixDSN, "")
		dialector = mysql.Open(opt.DSN)
	} else if strings.HasPrefix(opt.DSN, PostgresPrefixDSN) {
		dialector = postgres.Open(opt.DSN)
	} else {
		return nil, gorm.ErrUnsupportedDriver
	}

	db, err := gorm.Open(dialector, gormConf)
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err = sqlDB.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
