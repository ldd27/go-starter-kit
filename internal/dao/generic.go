package dao

import (
	"github.com/ldd27/go-starter-kit/internal/dao/builder"
	"github.com/ldd27/go-starter-kit/internal/model"
	"gorm.io/gorm"
)

type ExampleDao = DaoT[model.Example, *builder.ExampleBuilder]

func NewExampleDao(db *gorm.DB) DaoT[model.Example, *builder.ExampleBuilder] {
	return NewDaoT[model.Example, *builder.ExampleBuilder](db)
}
