package model

import (
	"time"
)

type IntModel struct {
	ID        int `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Int64Model struct {
	ID        int64 `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type StringModel struct {
	ID        string `gorm:"type:varchar(36);primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
