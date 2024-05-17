package g

import (
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	dataDB *gorm.DB
)

func SetDB(_db *gorm.DB) {
	db = _db
}

func DB() *gorm.DB {
	return db
}

func SetDataDB(_db *gorm.DB) {
	dataDB = _db
}

func DataDB() *gorm.DB {
	return dataDB
}
