package dto

type User struct {
	Id string `json:"id,omitempty"`
	// Имя
	FirstName string `json:"first_name,omitempty"`
	// Фамилия
	SecondName string `json:"second_name,omitempty"`
	// Дата рождения
	Birthdate string `json:"birthdate,omitempty"`
	// Интересы
	Biography string `json:"biography,omitempty"`
	// Город
	City string `json:"city,omitempty"`
}
