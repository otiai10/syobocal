package models

import "github.com/jinzhu/gorm"

var _db gorm.DB

// Init ...
func Init(db gorm.DB) gorm.DB {
	_db = db

	db.AutoMigrate(&Anime{})
	return _db
}

// DB ...
func DB() gorm.DB {
	return _db
}
