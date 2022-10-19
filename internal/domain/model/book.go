package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Isbn       string `gorm:"index:idx_books_isbn_unique,unique"`
	Title      string
	Author     string
	Descrption string
	Users      []User `gorm:"many2many:users_books"`
}
