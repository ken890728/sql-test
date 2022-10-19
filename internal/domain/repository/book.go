package repository

import "sql-test/internal/domain/model"

type BookRepo interface {
	Create(book *model.Book) error
	GetByUserId(userId string, cmpBookData *model.Book) ([]model.Book, error)
}
