package models

type Auth struct {
	Salt       string `default:"mysalt" json:"salt"`
	SigningKey string `default:"UNSECURE_SIGNING_KEY_EXAMPLE" json:"signingkey"`
	TokenTTL   uint   `default:"86400" json:"tokenttl"`
}
