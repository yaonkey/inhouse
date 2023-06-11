package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Создает новый экземпляр базы данных,
// создавая при этом новое подключение
func Connect() (db *gorm.DB, err error) {
	db, err = gorm.Open(sqlite.Open(
		"inhouse.sqlite",
	), &gorm.Config{})
	return
}
