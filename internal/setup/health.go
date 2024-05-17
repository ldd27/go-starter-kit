package setup

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ldd27/go-starter-kit/internal/conf"
	"github.com/ldd27/go-starter-kit/pkg/echox"
)

func SetupHealth(conf conf.Conf) error {
	r := echox.New()
	r.Group("").GET("/health", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "ok")
	})

	go func() {
		err := r.Start(fmt.Sprintf(":%d", conf.Server.Port))
		if err != nil {
			panic(err)
		}
	}()
	return nil
}
