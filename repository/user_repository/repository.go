package user_repository

import (
	"chapter3-sesi2/entity"
	"chapter3-sesi2/pkg/errs"
)

type UserRepository interface {
	CreateNewUser(user entity.User) errs.MessageErr
	GetUserById(userId int) (*entity.User, errs.MessageErr)
	GetUserByEmail(userEmail string) (*entity.User, errs.MessageErr)
}
