package dao

import (
	"context"

	"github.com/ldd27/go-starter-kit/internal/dao/builder"
	"gorm.io/gorm"
)

type DaoT[T interface{}, TBuilder builder.BuilderT] struct {
	db *gorm.DB
}

func NewDaoT[T interface{}, TBuilder builder.BuilderT](db *gorm.DB) DaoT[T, TBuilder] {
	return DaoT[T, TBuilder]{db: db}
}

func (r *DaoT[T, TBuilder]) Build(ctx context.Context, builder builder.BuilderT) *gorm.DB {
	if builder.Unscoped() {
		return r.db.Model(new(T)).WithContext(ctx).Unscoped().Clauses(builder.Build()...)
	} else {
		return r.db.Model(new(T)).WithContext(ctx).Clauses(builder.Build()...)
	}
}

func (r *DaoT[T, TBuilder]) BuildPage(ctx context.Context, builder builder.BuilderT, pageIndex, pageSize int) *gorm.DB {
	if builder.Unscoped() {
		return r.db.Model(new(T)).WithContext(ctx).Unscoped().Clauses(builder.Build()...).Offset((pageIndex - 1) * pageSize).Limit(pageSize)
	} else {
		return r.db.Model(new(T)).WithContext(ctx).Clauses(builder.Build()...).Offset((pageIndex - 1) * pageSize).Limit(pageSize)
	}
}

func (r *DaoT[T, TBuilder]) BuildCursorPage(ctx context.Context, builder builder.BuilderT, pageSize int) *gorm.DB {
	if builder.Unscoped() {
		return r.db.Model(new(T)).WithContext(ctx).Unscoped().Clauses(builder.Build()...).Limit(pageSize)
	} else {
		return r.db.Model(new(T)).WithContext(ctx).Clauses(builder.Build()...).Limit(pageSize)
	}
}

func (r *DaoT[T, TBuilder]) BuildDelete(ctx context.Context, builder builder.BuilderT) *gorm.DB {
	return r.db.Model(new(T)).WithContext(ctx).Clauses(builder.Build()...).Delete(new(T))
}

func (r *DaoT[T, TBuilder]) Ctx(ctx context.Context) *gorm.DB {
	return r.db.Model(new(T)).WithContext(ctx)
}

func (r *DaoT[T, TBuilder]) GetOneT(ctx context.Context, builder TBuilder) (result T, err error) {
	err = r.Build(ctx, builder).Take(&result).Error
	return
}

func (r *DaoT[T, TBuilder]) GetListT(ctx context.Context, builder TBuilder) (result []T, err error) {
	err = r.Build(ctx, builder).Find(&result).Error
	return
}

func (r *DaoT[T, TBuilder]) GetPageListT(ctx context.Context, builder TBuilder, pageIndex, pageSize int) (
	result []T, total int64, err error) {
	if err = r.BuildPage(ctx, builder, pageIndex, pageSize).Find(&result).Error; err != nil {
		return
	}
	if err = r.Build(ctx, builder).Count(&total).Error; err != nil {
		return
	}
	return
}

func (r *DaoT[T, TBuilder]) GetCursorListT(ctx context.Context, builder TBuilder, pageSize int) (
	result []T, err error) {
	if err = r.BuildCursorPage(ctx, builder, pageSize).Find(&result).Error; err != nil {
		return
	}
	return
}

func (r *DaoT[T, TBuilder]) CreateT(ctx context.Context, item T) (T, error) {
	if err := r.db.WithContext(ctx).Create(&item).Error; err != nil {
		return item, err
	}
	return item, nil
}

func (r *DaoT[T, TBuilder]) BatchCreateT(items []T) ([]T, error) {
	if err := r.db.Create(&items).Error; err != nil {
		return items, err
	}
	return items, nil
}

func (r *DaoT[T, TBuilder]) UpdateT(ctx context.Context, builder TBuilder, upt map[string]interface{}) (err error) {
	err = r.Build(ctx, builder).Updates(upt).Error
	return
}

func (r *DaoT[T, TBuilder]) UpdateWithRowsAffectedT(ctx context.Context, builder TBuilder, upt map[string]interface{}) (rowsAffected int64, err error) {
	res := r.Build(ctx, builder).Updates(upt)
	return res.RowsAffected, res.Error
}

func (r *DaoT[T, TBuilder]) DeleteT(ctx context.Context, builder TBuilder) (err error) {
	err = r.BuildDelete(ctx, builder).Error
	return err
}

func (r *DaoT[T, TBuilder]) DeleteWithRowsAffectedT(ctx context.Context, builder TBuilder) (
	rowsAffected int64, err error) {
	res := r.BuildDelete(ctx, builder)
	return res.RowsAffected, res.Error
}
