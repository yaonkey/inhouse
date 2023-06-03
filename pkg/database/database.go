package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

// Создает новый экземпляр базы данных,
// создавая при этом новое подключение
func Connect() (err error) {
	db, err = gorm.Open(sqlite.Open(
		"inhouse.sqlite",
	), &gorm.Config{})
	return
}

// Возвращает экземпляр базы данных
func GetDB() *gorm.DB {
	return db
}
