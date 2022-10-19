package repository

import (
	"log"
	"sql-test/internal/domain/model"
	"sql-test/internal/domain/repository"

	"gorm.io/gorm"
)

type mysqlBookRepo struct {
	DB *gorm.DB
}

func NewMysqlBookRepo(DB *gorm.DB) repository.BookRepo {
	DB.AutoMigrate()
	return &mysqlBookRepo{DB: DB}
}

func (bookRepo *mysqlBookRepo) Create(book *model.Book) error {
	err := bookRepo.DB.Save(book).Error
	if err != nil {
		log.Print("Repository error:", err)
		return err
	}
	return nil
}

func (bookRepo *mysqlBookRepo) GetByUserId(userId string) ([]model.Book, error) {
	var Books []model.Book
	user := model.User{UserId: userId}
	err := bookRepo.DB.Where(&model.Book{Users: []model.User{user}}).Find(&Books).Error
	return Books, err
}
