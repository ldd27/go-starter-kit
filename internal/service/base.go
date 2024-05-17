package service

import (
	"context"

	"github.com/ldd27/go-starter-kit/pkg/zapx"
	"go.uber.org/zap"
)

type service struct{}

func (r *service) Logger(ctx context.Context) *zap.Logger {
	return zapx.Logger(ctx)
}
