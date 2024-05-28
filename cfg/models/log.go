package models

type Log struct {
	File       string `default:"app.log" json:"file"`
	Level      string `default:"ERROR" json:"level"`
	Http       bool   `default:"false" json:"http"`
	Stack      bool   `default:"false" json:"stack"`
	StdOut     bool   `default:"true" json:"stdOut"`
	MaxSize    int    `default:"100" json:"maxSize"`
	MaxAge     int    `default:"28" json:"maxAge"`
	MaxBackups int    `default:"5" json:"maxBackups"`
	Compress   bool   `default:"true" json:"compress"`
}
