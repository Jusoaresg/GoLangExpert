package database

import "github.com/Jusoaresg/GoLangExpert/7-APIs/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	GetByEmail(emaild string) (*entity.User, error)
}
