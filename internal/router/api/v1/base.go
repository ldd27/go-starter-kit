package v1

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ldd27/go-starter-kit/internal/types"
	"github.com/ldd27/go-starter-kit/pkg/echox"
	"github.com/ldd27/go-starter-kit/pkg/zapx"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type controller struct {
	echoCtx echo.Context
	ctx     context.Context
	logger  *zap.Logger
}

func parseCtx(c echo.Context) controller {
	ctx := echox.Ctx(c)
	return controller{c, ctx, zapx.Logger(ctx)}
}

func (r *controller) BindValidate(req interface{}) error {
	if err := r.echoCtx.Bind(req); err != nil {
		r.logger.WithOptions(zap.AddCallerSkip(1)).Error("bind err", zap.Error(err))
		return err
	}

	if err := r.echoCtx.Validate(req); err != nil {
		r.logger.WithOptions(zap.AddCallerSkip(1)).Error("validate err", zap.Error(err))
		return err
	}
	return nil
}

func (r *controller) JsonRes(i interface{}) error {
	res := types.Res{
		Success: true,
		Data:    i,
	}
	return r.echoCtx.JSON(http.StatusOK, res)
}

func (r *controller) JsonPageRes(total int64, i interface{}) error {
	res := types.Res{
		Success: true,
		Data: types.PageRes{
			Data:  i,
			Total: total,
		},
	}
	return r.echoCtx.JSON(http.StatusOK, res)
}

func pkgCursorRes[T any](cursor int, items []T, fn func(item T) int) types.CursorPageRes {
	res := types.CursorPageRes{
		Data: items,
	}

	if len(items) == 0 {
		res.NextCursor = cursor
	} else {
		last, _ := lo.Last(items)
		res.NextCursor = fn(last)
	}
	return res
}
