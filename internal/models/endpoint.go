package models

import "yaonkey/inhouse/internal/database"

// Модель эндпоита, содержащая url
// и общее количество запросов count к эндпоинту
type Endpoint struct {
	Url   string `gorm:"primarykey;unique" json:"endpoint_url"`
	Count int    `json:"endpoint_count"`
}

// Модель множества эндпоинтов
type Endpoints []Endpoint

// Позволяет инкрементировать количество запросов к эндпоинту
func (e *Endpoint) Increment() (err error) {
	db, err := database.Connect()
	if err != nil {
		return
	}
	db.First(&e)
	e.Count++
	db.Save(&e)
	return
}

// Позволяет получить список всех эндпоинтов
func GetAllEndpoints() (*Endpoints, error) {
	endpoints := &Endpoints{}
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	db.Find(&endpoints)
	return endpoints, nil
}
