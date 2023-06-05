package models

import (
	"time"
	"yaonkey/inhouse/pkg/database"
)

// Модель сайта, содержащая url сайта
// и время time доступа к нему
type Site struct {
	Url  string        `gorm:"primarykey;unique" json:"sitename"`
	Time time.Duration `json:"sitetime"`
}

// Модель множества сайтов sites
type Sites []Site

// Позволяет записать время в модель сайта
func (s *Site) GetTime() (err error) {
	db, err := database.Connect()
	if err != nil {
		return
	}
	err = db.First(&s).Error
	return
}

// Позволяет записать сайт с минимальным временем доступа к нему
func (s *Site) GetMinTime() (err error) {
	db, err := database.Connect()
	db.Raw("SELECT * FROM sites WHERE time = (SELECT MIN(time) FROM sites)").Scan(&s)

	return
}

// Позволяет записать сайт с максимальным временем доступа к нему
func (s *Site) GetMaxTime() (err error) {
	db, err := database.Connect()
	db.Raw("SELECT * FROM sites WHERE time = (SELECT MAX(time) FROM sites)").Scan(&s)

	return
}
