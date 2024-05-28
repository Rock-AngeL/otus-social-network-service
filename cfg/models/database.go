package models

type Database struct {
	Dialect        string `json:"dialect" default:"postgres"`
	Host           string `json:"host" default:"10.0.4.91"`
	Port           string `json:"port" default:"5432"`
	User           string `json:"user" default:"postgres"`
	Password       string `json:"password" default:"geolook"`
	Name           string `json:"name" default:"socialnetwork"`
	SslMode        string `json:"sslmode" default:"disable"`
	Recreate       bool   `json:"recreate" default:"false"`
	InitPrivileges bool   `json:"initPrivileges" default:"false"`
	Defaults       bool   `json:"defaults" default:"false"`
	Timeout        int    `json:"timeout" default:"-1"`
	Prefix         string `json:"prefix" default:""`
}
