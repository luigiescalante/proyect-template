package database

import (
	"github.com/luigiescalante/proyect-template/core/domain"
	"github.com/luigiescalante/proyect-template/infrastructure/repository/database"
)

type repositoryUsers struct {
}

func (repo repositoryUsers) GetById(id int) (*domain.Users, error) {
	var user *domain.Users
	tx := database.Engine().Where("id=?", id).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func UserRepositoryFactory() *repositoryUsers {
	return &repositoryUsers{}
}

func (repo repositoryUsers) Save(user *domain.Users) error {
	tx := database.Engine().Save(user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo repositoryUsers) Delete(user *domain.Users) error {
	tx := database.Engine().Delete(user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
