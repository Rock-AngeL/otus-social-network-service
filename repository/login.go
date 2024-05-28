package repository

import (
	"otus-social-network-service_gen_swagger/db"
)

type LoginRepositoryInstance struct {
	dm *db.DataManager
	/* tokenTTL   uint
	signingKey string
	salt       string */
}

func NewLoginRepository(dm *db.DataManager) *LoginRepositoryInstance {
	return &LoginRepositoryInstance{
		dm: dm,
		/* salt:       cfg.Config().Services.Service.Auth.Salt,
		signingKey: cfg.Config().Services.Service.Auth.SigningKey,
		tokenTTL:   cfg.Config().Services.Service.Auth.TokenTTL, */
	}
}
