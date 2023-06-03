package models

// Модель сайта, имплементирующая string
type Site struct {
	Url string `gorm:"primarykey;unique"`
}

// Множество моделей сайтов
type Sites []Site

// Функция, инициирующая url сайтов, требующихся для работы
func InitSites() (sites *Sites) {

	return
}
