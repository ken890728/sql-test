package usecase

import (
	"sql-test/internal/domain/model"
	"sql-test/internal/domain/model/jwt"
	"sql-test/internal/domain/repository"
	"sql-test/internal/domain/usecase"
)

type userUsecase struct {
	repo     repository.UserRepo
	bookRepo repository.BookRepo
}

func NewUserUsecase(repo repository.UserRepo, bookRepo repository.BookRepo) usecase.UserUsecase {
	return &userUsecase{
		repo:     repo,
		bookRepo: bookRepo,
	}
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

func (userUsecase *userUsecase) UserLogin(userId string, passwd string) (usecase.JwtToken, error) {
	userData, err := userUsecase.repo.GetByUserId(userId)
	if err != nil {
		return "", err
	}
	err = userData.VerifyPasswd(passwd)
	if err != nil {
		return "", err
	}
	tokenStr, err := jwt.NewToken(userId)
	if err != nil {
		return "", err
	}
	return usecase.JwtToken(tokenStr), nil
}

func (userUsecase *userUsecase) UserData(userId string) (model.User, error) {
	userData, err := userUsecase.repo.GetByUserId(userId)
	if err != nil {
		return userData, err
	}
	return userData, nil
}

func (userUsecase *userUsecase) UserAddBook(userId string, book *model.Book) error {
	userData, err := userUsecase.repo.GetByUserId(userId)
	if err != nil {
		return nil
	}
	book.Users = append(book.Users, userData)
	err = userUsecase.bookRepo.Create(book)
	if err != nil {
		return err
	}
	return nil
}

func (userUsecase *userUsecase) UserGetBooks(userId string, cmpBookData *model.Book) ([]model.Book, error) {
	bookData, err := userUsecase.bookRepo.GetByUserId(userId, cmpBookData)
	if err != nil {
		return bookData, err
	}
	return bookData, nil
}
