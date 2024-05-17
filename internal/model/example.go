package model

import (
	"gorm.io/plugin/soft_delete"
)

type Example struct {
	IntModel
	Name      string                `gorm:"type:varchar(50);not null;default:'';"`
	DeletedAt soft_delete.DeletedAt `gorm:"softDelete:milli;not null;default:0;"`
}

func (r *Example) TableName() string {
	return "example"
}
