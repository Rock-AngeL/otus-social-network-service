package models

import "fmt"

type WebServer struct {
	Port int `default:"8080" json:"port"`
}

func (this *WebServer) ListenPort() string {
	return fmt.Sprintf(":%v", this.Port)
}
