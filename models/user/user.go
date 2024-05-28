package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id         uuid.UUID `json:"id" db:"id"`
	FirstName  string    `json:"first_name" db:"first_name" binding:"required,alphanum"`
	SecondName string    `json:"second_name" db:"second_name" binding:"required,alphanum"`
	Email      string    `json:"email" binding:"required,email"`
	Birthdate  string    `json:"birthdate" binding:"required"`
	Biography  string    `json:"biography" binding:"alphanum" faker:"lang=rus, sentence, "`
	City       string    `json:"city" binding:"alphanum" faker:"city, lang=rus"`
	Password   string    `json:"-" binding:"required,alphanum" faker:"word, lang=rus"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}
