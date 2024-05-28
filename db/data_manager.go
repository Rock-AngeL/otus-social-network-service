package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"otus-social-network-service_gen_swagger/cfg"
)

type DataManager struct {
	db *sqlx.DB
}

func (h *DataManager) Db() *sqlx.DB {
	return h.db
}

func (h *DataManager) SetDb(db *sqlx.DB) {
	h.db = db
}

func NewDataManager() *DataManager {
	return &DataManager{}
}

func (h *DataManager) Init() error {
	connectString := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=%v",
		cfg.Config().Services.Service.Database.Host,
		cfg.Config().Services.Service.Database.Port,
		cfg.Config().Services.Service.Database.User,
		cfg.Config().Services.Service.Database.Password,
		cfg.Config().Services.Service.Database.SslMode)
	db, err := sqlx.Connect(cfg.Config().Services.Service.Database.Dialect, connectString)

	_, err = db.Exec(fmt.Sprintf("drop database if exists %s", cfg.Config().Services.Service.Database.Name))
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = db.Exec(fmt.Sprintf("create database %s", cfg.Config().Services.Service.Database.Name))
	if err != nil {
		fmt.Println(err.Error())
	}
	db.Close()

	connectString = fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		cfg.Config().Services.Service.Database.Host,
		cfg.Config().Services.Service.Database.Port,
		cfg.Config().Services.Service.Database.Name,
		cfg.Config().Services.Service.Database.User,
		cfg.Config().Services.Service.Database.Password,
		cfg.Config().Services.Service.Database.SslMode)
	db, err = sqlx.Connect(cfg.Config().Services.Service.Database.Dialect, connectString)

	h.db = db

	h.createUserTable()

	return err
}

func (h *DataManager) createUserTable() error {
	query := "CREATE TABLE IF NOT EXISTS users (id uuid DEFAULT gen_random_uuid() PRIMARY KEY, first_name varchar(128) NOT NULL, second_name varchar(128) NOT NULL, birthdate TEXT NULL, city varchar(128) NOT NULL, biography TEXT, email varchar(128) NOT NULL UNIQUE, password TEXT NOT NULL, created_at timestamp, updated_at timestamp)"

	_, err := h.db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
	}

	return err
}
