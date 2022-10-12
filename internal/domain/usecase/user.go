package usecase

import "sql-test/internal/domain/model"

type JwtToken string

type UserUsecase interface {
	UserRegister(user *model.User) error
	UserDelete(ID int) error
	UserLogin(userId string, passwd string) (JwtToken, error)
	UserData(userId string) (model.User, error)
}
