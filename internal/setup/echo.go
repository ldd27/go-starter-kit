package setup

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/ldd27/go-starter-kit/internal/conf"
	"github.com/ldd27/go-starter-kit/internal/constant"
	"github.com/ldd27/go-starter-kit/internal/router"
	"github.com/ldd27/go-starter-kit/internal/setup/g"
	"github.com/ldd27/go-starter-kit/internal/types"
	"github.com/ldd27/go-starter-kit/pkg/echox"
	"github.com/ldd27/go-starter-kit/pkg/zapx"
	"go.uber.org/zap"
)

func SetupEcho(conf conf.Conf, wg *sync.WaitGroup, cancel context.CancelFunc) error {
	engine := echox.New(func(opt *echox.Option) {
		opt.Metrics.Enabled = true
		opt.Pprof.Enabled = true
		opt.CORS.Enabled = conf.Debug
		opt.ErrorHandler = EchoErrorHandler
	})

	if conf.Server.EnableSwagger {
		engine.File("/swagger.yaml", "./docs/swagger.yaml")
	}

	router.InitRouter(engine)

	logger := zap.L()

	idleConnsClosed := make(chan struct{})
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		logger.Info(fmt.Sprintf("Received signal %s", <-ch))

		cancel()

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		if err := engine.Shutdown(ctx); err != nil {
			logger.Error("shutdown err", zap.Error(err))
		}
		logger.Info("server shutdown success")

		close(idleConnsClosed)
	}()

	logger.Info("server init success")
	if err := engine.Start(fmt.Sprintf(":%d", 8080)); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Error("run server err", zap.Error(err))
		return err
	}

	<-idleConnsClosed
	return nil
}

func EchoErrorHandler(err error, c echo.Context) {
	res := NewErrRes(err)

	if err = c.JSON(res.StatusCode, res); err != nil {
		zapx.Logger(echox.Ctx(c)).Error("json err", zap.Error(err))
		return
	}
}

func NewErrRes(err error) types.Res {
	var res types.Res

	var errCode constant.ErrCode
	var validationErrors validator.ValidationErrors
	var httpError *echo.HTTPError

	if errors.As(err, &errCode) {
		if errors.Is(errCode, constant.ErrUnauthorized) {
			res = types.NewErrResWithErrCode(http.StatusUnauthorized, errCode)
		} else {
			res = types.NewErrResWithErrCode(http.StatusOK, errCode)
		}
	} else if errors.As(err, &validationErrors) {
		res = types.NewErrResWithStatusCode(http.StatusBadRequest, constant.ErrInvalidParams.ErrCode, validationErrors.Error())
	} else if errors.As(err, &httpError) {
		res = types.NewErrResWithStatusCode(httpError.Code, httpError.Code, httpError.Error())
	} else if g.C.Debug {
		res = types.NewErrResWithStatusCode(http.StatusInternalServerError, constant.ErrInternal.ErrCode, err.Error())
	} else {
		res = types.NewErrResWithErrCode(http.StatusInternalServerError, constant.ErrInternal)
	}

	return res
}
