//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/ldd27/go-starter-kit/internal/dao"
	v1 "github.com/ldd27/go-starter-kit/internal/router/api/v1"
	"github.com/ldd27/go-starter-kit/internal/service"
	"github.com/ldd27/go-starter-kit/internal/setup/g"
)

func NewExampleController() *v1.ExampleController {
	wire.Build(g.DB, dao.NewExampleDao, service.NewExampleSvc, v1.NewExampleController)
	return &v1.ExampleController{}
}
