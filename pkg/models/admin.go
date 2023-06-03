package models

// Модель администратора, содержащая
// информацию о логине и пароле администратора
type Admin struct {
	Login         string `gorm:"primarykey;unique"`
	Password_hash string
}
