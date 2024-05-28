package dto

type UserRegisterBody struct {
	FirstName string `json:"first_name,omitempty"`

	SecondName string `json:"second_name,omitempty"`

	Birthdate string `json:"birthdate,omitempty"`

	Biography string `json:"biography,omitempty"`

	City string `json:"city,omitempty"`

	Email string `json:"email,omitempty"`

	Password string `json:"password,omitempty"`
}
