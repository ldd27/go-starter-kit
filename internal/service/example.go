package service

import (
	"context"

	"github.com/ldd27/go-starter-kit/internal/constant"
	"github.com/ldd27/go-starter-kit/internal/dao"
	"github.com/ldd27/go-starter-kit/internal/dao/builder"
	"github.com/ldd27/go-starter-kit/internal/model"
	"go.uber.org/zap"
)

type ExampleSvc struct {
	service
	exampleDao dao.ExampleDao
}

func NewExampleSvc(exampleDao dao.ExampleDao) *ExampleSvc {
	return &ExampleSvc{
		exampleDao: exampleDao,
	}
}

func (r *ExampleSvc) GetList(ctx context.Context) (result []model.Example, err error) {
	scope := builder.NewExampleBuilder().OrderByIDDesc()
	result, err = r.exampleDao.GetListT(ctx, scope)
	if err != nil {
		r.Logger(ctx).Error("get example list", zap.Error(err))
		return nil, constant.ErrInternal
	}
	return result, nil
}

func (r *ExampleSvc) GetPageList(ctx context.Context, pageIndex, pageSize int) (result []model.Example, total int64, err error) {
	scope := builder.NewExampleBuilder().OrderByIDDesc()
	result, total, err = r.exampleDao.GetPageListT(ctx, scope, pageIndex, pageSize)
	if err != nil {
		r.Logger(ctx).Error("get example list", zap.Error(err))
		return nil, 0, constant.ErrInternal
	}
	return result, total, nil
}

func (r *ExampleSvc) GetCursorPageList(ctx context.Context, cursor, pageSize int) (result []model.Example, err error) {
	scope := builder.NewExampleBuilder().WithCursor(cursor)
	result, err = r.exampleDao.GetCursorListT(ctx, scope, pageSize)
	if err != nil {
		r.Logger(ctx).Error("get example list", zap.Error(err))
		return nil, constant.ErrInternal
	}
	return result, nil
}

func (r *ExampleSvc) Create(ctx context.Context, name string) (err error) {
	item := model.Example{
		Name: name,
	}
	if _, err = r.exampleDao.CreateT(ctx, item); err != nil {
		r.Logger(ctx).Error("create example err", zap.Error(err))
		return constant.ErrInternal
	}
	return nil
}

func (r *ExampleSvc) Update(ctx context.Context, id int, name string) (err error) {
	scope := builder.NewExampleBuilder().WithID(id)
	upt := map[string]interface{}{
		"name": name,
	}
	if err = r.exampleDao.UpdateT(ctx, scope, upt); err != nil {
		r.Logger(ctx).Error("update example err", zap.Error(err))
		return constant.ErrInternal
	}
	return nil
}

func (r *ExampleSvc) Delete(ctx context.Context, id int) (err error) {
	scope := builder.NewExampleBuilder().WithID(id)
	if err = r.exampleDao.DeleteT(ctx, scope); err != nil {
		r.Logger(ctx).Error("delete example err", zap.Error(err))
		return constant.ErrInternal
	}
	return nil
}
