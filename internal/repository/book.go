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

func (bookRepo *mysqlBookRepo) GetByUserId(userId string, cmpBookData *model.Book) ([]model.Book, error) {
	var user model.User
	//err := bookRepo.DB.Where(&model.Book{Users: []model.User{user}}).Find(&Books).Error
	cmp := bookRepo.DB
	if cmpBookData.Isbn != "" {
		cmp = cmp.Where("books.isbn = ?", cmpBookData.Isbn)
	}
	if cmpBookData.Title != "" {
		cmp = cmp.Where("books.title LIKE ?", "%"+cmpBookData.Title+"%")
	}
	if cmpBookData.Author != "" {
		cmp = cmp.Where("books.author LIKE ?", "%"+cmpBookData.Author+"%")
	}

	bookRepo.DB.Preload("Books", cmp).
		Where("users.user_id = ?", userId).
		Find(&user)

	return user.Books, nil
}
