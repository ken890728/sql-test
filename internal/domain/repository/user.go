package repository

import "sql-test/internal/domain/model"

type UserRepo interface {
	Create(user *model.User) error
	Delete(ID int) error
}
