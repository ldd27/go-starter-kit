package gormx

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ldd27/go-starter-kit/pkg/zapx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type CustomLogger struct {
	LogLevel                            logger.LogLevel
	traceStr, traceErrStr, traceWarnStr string
	SlowThreshold                       time.Duration
	IgnoreRecordNotFoundError           bool
	Colorful                            bool
}

func NewLogger(opt Option) logger.Interface {
	var (
		traceStr     = "[%.3fms] [rows:%v] %s"
		traceWarnStr = "%s [%.3fms] [rows:%v] %s"
		traceErrStr  = "%s [%.3fms] [rows:%v] %s"
	)

	if opt.Colorful {
		traceStr = logger.Yellow + "[%.3fms] " + logger.BlueBold + "[rows:%v]" + logger.Reset + " %s"
		traceWarnStr = logger.Green + "%s " + logger.Reset + logger.RedBold + "[%.3fms] " + logger.Yellow + "[rows:%v]" + logger.Magenta + " %s" + logger.Reset
		traceErrStr = logger.MagentaBold + "%s " + logger.Reset + logger.Yellow + "[%.3fms] " + logger.BlueBold + "[rows:%v]" + logger.Reset + " %s"
	}

	return &CustomLogger{
		traceStr:      traceStr,
		traceWarnStr:  traceWarnStr,
		traceErrStr:   traceErrStr,
		SlowThreshold: time.Duration(opt.SlowThreshold) * time.Millisecond,
	}
}

func (l *CustomLogger) LogMode(level logger.LogLevel) logger.Interface {
	l.LogLevel = level
	return l
}

func (l *CustomLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		zapx.Logger(ctx).WithOptions(zap.AddCallerSkip(1)).Sugar().Info(append([]interface{}{msg}, data...)...)
	}
}

func (l *CustomLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		zapx.Logger(ctx).WithOptions(zap.AddCallerSkip(1)).Sugar().Warn(append([]interface{}{msg}, data...)...)
	}
}

func (l *CustomLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		zapx.Logger(ctx).WithOptions(zap.AddCallerSkip(1)).Sugar().Error(append([]interface{}{msg}, data...)...)
	}
}

func (l *CustomLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= logger.Error && (!errors.Is(err, gorm.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			l.Printf(ctx, zap.ErrorLevel, l.traceErrStr, err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.Printf(ctx, zap.ErrorLevel, l.traceErrStr, err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		if rows == -1 {
			l.Printf(ctx, zap.WarnLevel, l.traceWarnStr, slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.Printf(ctx, zap.WarnLevel, l.traceWarnStr, slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case l.LogLevel == logger.Info:
		sql, rows := fc()
		if rows == -1 {
			l.Printf(ctx, zap.InfoLevel, l.traceStr, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.Printf(ctx, zap.InfoLevel, l.traceStr, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}

func (l *CustomLogger) Printf(ctx context.Context, level zapcore.Level, format string, v ...interface{}) {
	switch level {
	case zap.ErrorLevel:
		zapx.Logger(ctx).WithOptions(zap.AddCallerSkip(4)).Sugar().Errorf(format, v...)
	case zap.WarnLevel:
		zapx.Logger(ctx).WithOptions(zap.AddCallerSkip(4)).Sugar().Warnf(format, v...)
	default:
		zapx.Logger(ctx).WithOptions(zap.AddCallerSkip(4)).Sugar().Infof(format, v...)
	}
}
