package usecase

import "sql-test/internal/domain/model"

type UserUsecase interface {
	UserRegister(user *model.User) error
	UserDelete(ID int) error
}
