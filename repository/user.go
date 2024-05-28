package repository

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"otus-social-network-service_gen_swagger/app_error"
	"otus-social-network-service_gen_swagger/db"
	dto "otus-social-network-service_gen_swagger/dto/user"
	model "otus-social-network-service_gen_swagger/models/user"
	"strings"
	"time"
)

type UserRepositoryInstance struct {
	dm *db.DataManager
}

func NewUserRepository(dm *db.DataManager) *UserRepositoryInstance {
	return &UserRepositoryInstance{dm: dm}
}

func (r *UserRepositoryInstance) CreateUser(user *dto.UserRegisterBody) (uuid.UUID, error) {
	rows, err := r.dm.Db().Query("SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)", user.Email)
	if err != nil {
		return uuid.Nil, app_error.NewInternalServerError(err)
	}

	var exists bool

	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&exists); err != nil {
			return uuid.Nil, app_error.NewInternalServerError(err)
		}

		if exists {
			return uuid.Nil, app_error.NewHttpError(err, fmt.Sprintf("User with email %s already registered", user.Email), "email", http.StatusBadRequest)
		}
	}

	query := "INSERT INTO users (first_name, second_name, email, biography, birthdate, city, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id"

	now := time.Now()
	var userId uuid.UUID
	err = r.dm.Db().QueryRow(query, user.FirstName, user.SecondName, user.Email, user.Biography, user.Birthdate, user.City, user.Password, now, now).Scan(&userId)

	if err != nil {
		return uuid.Nil, app_error.NewInternalServerError(err)
	}

	return userId, nil
}

func (r *UserRepositoryInstance) GetUser(email, password string) (*model.User, *app_error.HttpError) {
	rows, err := r.dm.Db().Queryx("SELECT * FROM users WHERE email=$1 and password=$2", email, password)
	if err != nil {
		return new(model.User), app_error.NewInternalServerError(err)
	}
	defer rows.Close()

	var user model.User
	for rows.Next() {
		err = rows.StructScan(&user)
		if err != nil {
			return new(model.User), app_error.NewInternalServerError(err)
		}
	}

	return &user, nil
}

func (r *UserRepositoryInstance) GetUserById(userId uuid.UUID) (*model.User, error) {
	var user model.User
	err := r.dm.Db().Get(&user, "SELECT * FROM users WHERE id=$1 LIMIT 1", userId)

	if err != nil {
		return new(model.User), app_error.NewHttpError(err, "user not found", "user_id", http.StatusBadRequest)
	}

	return &user, nil
}

func (r *UserRepositoryInstance) FindUsers(firstName, lastName string) ([]*model.User, error) {
	users := make([]*model.User, 100)
	query := "SELECT * FROM users WHERE "
	paramName := strings.ToLower(firstName) + "%"
	paramSurname := strings.ToLower(lastName) + "%"
	limitPart := " ORDER BY id LIMIT 100;"

	var err error

	if len(firstName) > 1 && len(lastName) > 1 {
		err = r.dm.Db().Select(&users, query+"(lower(first_name) LIKE $1 AND lower(second_name) LIKE $2) OR (lower(second_name) LIKE $3 AND lower(first_name) LIKE $4)"+limitPart, paramName, paramSurname, paramName, paramSurname)
	} else if len(firstName) > 0 {
		err = r.dm.Db().Select(&users, query+"lower(first_name) LIKE $1"+limitPart, paramName)
	} else if len(lastName) > 0 {
		err = r.dm.Db().Select(&users, query+"lower(second_name) LIKE $1"+limitPart, paramSurname)
	} else {
		err = r.dm.Db().Select(&users, query+limitPart)
	}

	if err != nil {
		return users, app_error.NewHttpError(err, "user not found", "users", http.StatusBadRequest)
	}

	return users, nil
}
