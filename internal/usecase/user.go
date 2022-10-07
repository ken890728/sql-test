package usecase

import (
	"sql-test/internal/domain/model"
	"sql-test/internal/domain/repository"
	"sql-test/internal/domain/usecase"
)

type userUsecase struct {
	repo repository.UserRepo
}

func NewUserUsecase(repo repository.UserRepo) usecase.UserUsecase {
	return &userUsecase{repo: repo}
}

func (userUsecase *userUsecase) UserRegister(user *model.User) error {
	user.HashPassword()
	err := userUsecase.repo.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (userUsecase *userUsecase) UserDelete(ID int) error {
	userUsecase.repo.Delete(ID)
	return nil
}
