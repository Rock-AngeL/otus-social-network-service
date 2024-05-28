package repository

import (
	"otus-social-network-service_gen_swagger/db"
)

type RepositoryManager struct {
	userRepositoryInstance  *UserRepositoryInstance
	loginRepositoryInstance *LoginRepositoryInstance
}

func NewRepositoryManager(dm *db.DataManager) *RepositoryManager {
	return &RepositoryManager{
		userRepositoryInstance:  NewUserRepository(dm),
		loginRepositoryInstance: NewLoginRepository(dm),
	}
}

func (r *RepositoryManager) UserRepositoryInstance() *UserRepositoryInstance {
	return r.userRepositoryInstance
}

func (r *RepositoryManager) LoginRepositoryInstance() *LoginRepositoryInstance {
	return r.loginRepositoryInstance
}
