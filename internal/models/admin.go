package models

import (
	"yaonkey/inhouse/internal/database"
)

// Модель администратора, содержащая
// информацию о логине login
// и хеше пароля password_hash администратора
type Admin struct {
	Login         string `gorm:"primarykey;unique" json:"admin_login"`
	Password_hash string `json:"admin_password_hash"`
}

// Позволяет запись данные об администраторе
func (a *Admin) GetAdminData() (err error) {
	db, err := database.Connect()
	if err != nil {
		return
	}
	err = db.First(&a).Error
	return
}

// Простая проверка на администратора
func IsAdmin(a *Admin) bool {
	db, err := database.Connect()
	if err != nil {
		return false
	}
	var u int64
	db.Model(&Admin{}).Where("login=? AND password_hash=?", a.Login, a.Password_hash).Count(&u)

	return u != 0
}
