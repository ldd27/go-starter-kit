package echox

import (
	"context"

	"github.com/brpaz/echozap"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type corsOption struct {
	Enabled bool
}

type pprofOption struct {
	Enabled bool
}

type metricsOption struct {
	Enabled bool
}

type Option struct {
	CORS         corsOption
	Pprof        pprofOption
	Metrics      metricsOption
	ErrorHandler echo.HTTPErrorHandler
}

var defaultOption = Option{}

type EchoValidator struct {
	validator *validator.Validate
}

func (cv *EchoValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func New(opts ...func(*Option)) *echo.Echo {
	opt := defaultOption
	for _, o := range opts {
		o(&opt)
	}

	e := echo.New()

	e.HideBanner = true

	validate := validator.New()
	// validate.SetTagName("v")
	e.Validator = &EchoValidator{validator: validate}

	// recovery
	e.Use(middleware.Recover())
	// request_id
	e.Use(middleware.RequestID())
	// logger
	e.Use(echozap.ZapLogger(zap.L()))

	// cors
	if opt.CORS.Enabled {
		e.Use(middleware.CORS())
	}

	// pprof
	// if opt.Pprof.Enabled {
	//	pprof.Register(e.engine)
	// }

	// metrics
	// if opt.Metrics.Enabled {
	//	p := prometheus.NewPrometheus("echo", nil)
	//	p.Use(e.engine)
	// }

	if opt.ErrorHandler != nil {
		e.HTTPErrorHandler = opt.ErrorHandler
	}

	return e
}

func Ctx(c echo.Context) context.Context {
	ctx := c.Request().Context()
	reqID := c.Response().Header().Get(echo.HeaderXRequestID)
	if reqID != "" {
		ctx = context.WithValue(ctx, echo.HeaderXRequestID, reqID)
	}
	return ctx
}

func NewCtx(c echo.Context) context.Context {
	ctx := context.Background()
	reqID := c.Response().Header().Get(echo.HeaderXRequestID)
	if reqID != "" {
		ctx = context.WithValue(ctx, echo.HeaderXRequestID, reqID)
	}
	return ctx
}
